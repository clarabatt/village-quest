package repository

import (
	"database/sql"
	"fmt"
	"villageQuest/domain/entity/game"
	"villageQuest/infra/database"

	"github.com/google/uuid"
)

type GameRepository interface {
	Insert(game.Game) error
	GetAll() ([]game.Game, error)
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

func (g *gameRepository) Insert(game game.Game) error {
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

func (g *gameRepository) GetAll() ([]game.Game, error) {
	query := `SELECT id, number, max_days_played, players_name FROM game`
	result := g.connection.Query(query)
	var games []game.Game

	if result.Err != nil {
		return nil, result.Err
	}
	defer result.Rows.Close()

	for result.Rows.Next() {
		var gameID uuid.UUID
		var number, maxDaysPlayed int
		var playersName string

		err := result.Rows.Scan(&gameID, &number, &maxDaysPlayed, &playersName)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		game := game.LoadGame(gameID, number, maxDaysPlayed, playersName)
		games = append(games, *game)
	}

	if len(games) == 0 {
		return nil, sql.ErrNoRows
	}

	return games, nil
}

func (g *gameRepository) GetNextGameNumber() (int, error) {
	query := `SELECT COALESCE(MAX(number), 0) + 1 FROM game`
	result := g.connection.Query(query)

	if result.Err != nil {
		return 0, result.Err
	}
	defer result.Rows.Close()

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
