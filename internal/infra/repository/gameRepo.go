package repository

import (
	"database/sql"
	"villageQuest/internal/domain/entity/game"
	"villageQuest/internal/infra/database"

	"github.com/google/uuid"
)

type GameRepository interface {
	Insert(game.Game) error
	GetById(uuid.UUID) (game.Game, error)
	GetByNumber(int) (game.Game, error)
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
	_, err := g.connection.Exec(query, game.Id(), game.Number(), game.MaxDaysPlayed(), game.PlayersName())
	if err != nil {
		return err
	}
	return nil
}

func (g *gameRepository) GetById(id uuid.UUID) (game.Game, error) {
	query := `SELECT * FROM game WHERE id = ?`
	result := g.connection.Query(query, id)

	if result.Err != nil {
		return game.Game{}, result.Err
	}
	defer result.Rows.Close()

	if result.Rows.Next() {
		var gameID uuid.UUID
		var number, maxDaysPlayed int
		var playersName string

		err := result.Rows.Scan(&gameID, &number, &maxDaysPlayed, &playersName)
		if err != nil {
			return game.Game{}, err
		}
		return *game.NewSavedGame(gameID, number, maxDaysPlayed, playersName), nil
	}

	return game.Game{}, sql.ErrNoRows
}

func (g *gameRepository) GetByNumber(num int) (game.Game, error) {
	query := `SELECT id, number, max_days_played, players_name FROM game WHERE number = ?`
	result := g.connection.Query(query, num)

	if result.Err != nil {
		return game.Game{}, result.Err
	}
	defer result.Rows.Close()

	if result.Rows.Next() {
		var gameID uuid.UUID
		var number, maxDaysPlayed int
		var playersName string

		err := result.Rows.Scan(&gameID, &number, &maxDaysPlayed, &playersName)
		if err != nil {
			return game.Game{}, err
		}

		return *game.NewSavedGame(gameID, number, maxDaysPlayed, playersName), nil
	}

	return game.Game{}, sql.ErrNoRows
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
