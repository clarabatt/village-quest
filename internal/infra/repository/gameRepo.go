package repository

import (
	"villageQuest/internal/domain/entity/game"
	"villageQuest/internal/infra/database"

	"github.com/google/uuid"
)

type GameRepository interface {
	Save(game.Game) error
	GetById(uuid.UUID) (game.Game, error)
	GetByNumber(int) (game.Game, error)
	GetNextGameNumber() (int, error)
}

type gameRepository struct {
	connection database.DatabaseConnection
}

func NewGameRepository(connection database.DatabaseConnection) GameRepository {
	return &gameRepository{
		connection: connection,
	}
}

func (g *gameRepository) Save(game game.Game) error {
	query := `
		INSERT INTO game (id, number, max_days_played, players_name)
		VALUES (?, ?, ?, ?)
	`
	_, err := g.connection.Query(query, game.Id(), game.Number(), game.MaxDaysPlayed(), game.PlayersName())
	return err
}

func (g *gameRepository) GetById(id uuid.UUID) (game.Game, error) {
	query := `SELECT * FROM game WHERE id = ?`
	row, err := g.connection.Query(query, id)

	var gameID uuid.UUID
	var number, maxDaysPlayed int
	var playersName string

	err = row.Scan(&gameID, &number, &maxDaysPlayed, &playersName)
	if err != nil {
		return game.Game{}, err
	}

	return *game.NewSavedGame(gameID, number, maxDaysPlayed, playersName), nil
}

func (g *gameRepository) GetByNumber(num int) (game.Game, error) {
	query := `SELECT id, number, max_days_played, players_name FROM game WHERE number = ?`
	row, err := g.connection.Query(query, num)

	var gameID uuid.UUID
	var number, maxDaysPlayed int
	var playersName string

	err = row.Scan(&gameID, &number, &maxDaysPlayed, &playersName)
	if err != nil {
		return game.Game{}, err
	}

	return *game.NewSavedGame(gameID, number, maxDaysPlayed, playersName), nil
}

func (g *gameRepository) GetNextGameNumber() (int, error) {
	var nextNumber int
	query := `SELECT COALESCE(MAX(number), 0) + 1 FROM game`
	row, err := g.connection.Query(query)

	err = row.Scan(&nextNumber)
	if err != nil {
		return 0, err
	}
	return nextNumber, nil
}
