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
	GetNextGameNumber() (int, error)
}

type gameRepository struct {
	connection database.DBAdapter
}

func NewGameRepository(connection database.DBAdapter) GameRepository {
	return &gameRepository{
		connection: connection,
	}
}

func (g *gameRepository) Insert(game Game) error {
	query := `
		INSERT INTO game (id, number, max_days_played, players_name)
		VALUES (?, ?, ?, ?)
	`
	_, err := g.connection.Exec(query, game.Id(), game.Number(), game.DaysPlayed(), game.PlayersName())
	if err != nil {
		return err
	}
	return nil
}

func (g *gameRepository) GetAll() ([]Game, error) {
	query := `SELECT id, number, max_days_played, players_name FROM game`
	result, err := g.connection.Query(query)
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

func (g *gameRepository) GetNextGameNumber() (int, error) {
	query := `SELECT COALESCE(MAX(number), 0) + 1 FROM game`
	result, err := g.connection.Query(query)
	if err != nil {
		return 0, err
	}
	defer func() {
		if err := result.Rows.Close(); err != nil {
			log.Printf("Warning: failed to close row: %v", err)
		}
	}()
	if result.Rows.Next() {
		var nextNumber int
		err := result.Rows.Scan(&nextNumber)
		if err != nil {
			return 0, err
		}
		return nextNumber, nil
	}

	return 0, sql.ErrNoRows
}
