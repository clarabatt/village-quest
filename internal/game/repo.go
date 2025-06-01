package game

import (
	"database/sql"
	"fmt"
	"log"
	"villagequest/internal/database"

	"github.com/google/uuid"
)

type GameRepository interface {
	Insert(Game) error
	GetAll() ([]Game, error)
	GetById(id string) (Game, error)
	Update(game Game) error
	Delete(id string) error
}

type gameRepository struct {
	connection database.DBAdapter
}

func NewGameRepository(connection database.DBAdapter) GameRepository {
	return &gameRepository{
		connection: connection,
	}
}

func (r *gameRepository) Insert(game Game) error {
	query := `
		INSERT INTO game (id, number, max_days_played, players_name)
		VALUES (?, ?, ?, ?)
	`
	_, err := r.connection.Execute(query, game.Id(), game.Number(), game.DaysPlayed(), game.PlayersName())
	if err != nil {
		return err
	}
	return nil
}

func (r *gameRepository) GetAll() ([]Game, error) {
	query := `SELECT id, number, max_days_played, players_name FROM game`
	result, err := r.connection.Query(query)
	if err != nil {
		return nil, err
	}
	var games []Game
	defer func() {
		if err := result.Rows.Close(); err != nil {
			log.Printf("Warning: failed to close row: %v", err)
		}
	}()

	for result.Rows.Next() {
		var gameID uuid.UUID
		var number, maxDaysPlayed int
		var playersName string

		err := result.Rows.Scan(&gameID, &number, &maxDaysPlayed, &playersName)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		game := LoadGame(gameID, number, maxDaysPlayed, playersName)
		games = append(games, *game)
	}

	if len(games) == 0 {
		return nil, sql.ErrNoRows
	}

	return games, nil
}

func (r *gameRepository) GetById(id string) (Game, error) {
	query := `SELECT id, number, max_days_played, players_name FROM game WHERE id = ?`
	result, err := r.connection.Query(query, id)
	if err != nil {
		return Game{}, err
	}
	defer result.Rows.Close()

	if !result.Rows.Next() {
		return Game{}, sql.ErrNoRows
	}

	var gameId, playersName string
	var number, maxDaysPlayed int

	err = result.Rows.Scan(&gameId, &number, &maxDaysPlayed, &playersName)
	if err != nil {
		return Game{}, err
	}

	return *LoadGame(uuid.MustParse(gameId), number, maxDaysPlayed, playersName), nil
}

func (r *gameRepository) Update(game Game) error {
	query := `
		UPDATE game 
		SET number = ?, max_days_played = ?, players_name = ?
		WHERE id = ?
	`
	_, err := r.connection.Execute(query, game.Number(), game.DaysPlayed(), game.PlayersName(), game.Id())
	return err
}

func (r *gameRepository) Delete(id string) error {
	query := `DELETE FROM game WHERE id = ?`
	_, err := r.connection.Execute(query, id)
	return err
}
