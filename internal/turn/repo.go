package turn

import (
	"database/sql"
	"villagequest/internal/database"

	"github.com/google/uuid"
)

type TurnRepository interface {
	Insert(Turn) error
	Update(Turn) (Turn, error)
	GetLastTurn() (Turn, error)
}

type turnRepository struct {
	connection database.DBAdapter
	gameId     uuid.UUID
}

func NewTurnRepository(connection database.DBAdapter, gameId uuid.UUID) TurnRepository {
	return &turnRepository{
		connection: connection,
		gameId:     gameId,
	}
}

func (r *turnRepository) Insert(turn Turn) error {
	query := `
		INSERT INTO turn_history (id, game_id, number, players_action, random_events, initial_resources, final_resources)
		VALUES(?, ?, ?, ?, ?, ?)
	`
	_, err := r.connection.Execute(
		query,
		turn.Id,
		r.gameId,
		turn.Number,
		turn.PlayersAction,
		turn.RandomEvents,
		turn.InitialResources,
		turn.FinalResources,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *turnRepository) Update(turn Turn) (Turn, error) {
	query := `
		UPDATE turn_history
		SET players_action = ?, random_events = ?, initial_resources = ?, final_resources = ?
		WHERE id = ? AND game_id = ?
	`
	_, err := r.connection.Execute(
		query,
		turn.PlayersAction,
		turn.RandomEvents,
		turn.InitialResources,
		turn.FinalResources,
		turn.Id,
		r.gameId,
	)
	if err != nil {
		return Turn{}, err
	}
	return turn, nil
}

func (r *turnRepository) GetLastTurn() (Turn, error) {
	query := `
		SELECT id, number, players_action, random_events, initial_resources, final_resources
		FROM turn_history
		WHERE game_id = ?
		ORDER BY number DESC
		LIMIT 1
	`
	result, err := r.connection.Query(query, r.gameId)
	if err != nil {
		return Turn{}, err
	}
	defer result.Rows.Close()

	var turn Turn
	if !result.Rows.Next() {
		return Turn{}, sql.ErrNoRows
	}
	err = result.Rows.Scan(&turn.Id, &turn.Number, &turn.PlayersAction, &turn.RandomEvents, &turn.InitialResources, &turn.FinalResources)
	if err != nil {
		return Turn{}, err
	}

	return turn, nil
}
