package turn

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type turnRepository struct {
	db *gorm.DB
}

type TurnRepository interface {
	Create(turn Turn, gameID uuid.UUID) (*Turn, error)
	Update(turn Turn, gameID uuid.UUID) (*Turn, error)
	GetLastTurn(gameID uuid.UUID) (*Turn, error)
	GetTurnByID(id uuid.UUID, gameID uuid.UUID) (*Turn, error)
	GetAllTurns(gameID uuid.UUID) ([]Turn, error)
}

func (TurnModel) TableName() string {
	return "turn_history"
}

func NewTurnRepository(db *gorm.DB) TurnRepository {
	return &turnRepository{
		db: db,
	}
}

func (r *turnRepository) Create(turn Turn, gameID uuid.UUID) (*Turn, error) {
	model, err := turnToModel(&turn, gameID)
	if err != nil {
		return nil, fmt.Errorf("failed to convert turn to model: %w", err)
	}

	if err := r.db.Create(model).Error; err != nil {
		return nil, fmt.Errorf("failed to insert turn: %w", err)
	}

	result, err := modelToTurn(model)
	if err != nil {
		return nil, fmt.Errorf("failed to convert model back to turn: %w", err)
	}

	return result, nil
}

func (r *turnRepository) Update(turn Turn, gameID uuid.UUID) (*Turn, error) {
	model, err := turnToModel(&turn, gameID)
	if err != nil {
		return nil, fmt.Errorf("failed to convert turn to model: %w", err)
	}

	result := r.db.Model(&TurnModel{}).
		Where("id = ? AND game_id = ?", turn.Id, gameID).
		Updates(model)

	if result.Error != nil {
		return nil, fmt.Errorf("failed to update turn: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("turn not found: id=%s, gameID=%s", turn.Id, gameID)
	}

	var updatedModel TurnModel
	if err := r.db.Where("id = ? AND game_id = ?", turn.Id, gameID).First(&updatedModel).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch updated turn: %w", err)
	}

	updatedTurn, err := modelToTurn(&updatedModel)
	if err != nil {
		return nil, fmt.Errorf("failed to convert updated model to turn: %w", err)
	}

	return updatedTurn, nil
}

func (r *turnRepository) GetLastTurn(gameID uuid.UUID) (*Turn, error) {
	var model TurnModel

	err := r.db.Where("game_id = ?", gameID).
		Order("number DESC").
		First(&model).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("no turns found for game %s", gameID)
		}
		return nil, fmt.Errorf("failed to get last turn: %w", err)
	}

	turn, err := modelToTurn(&model)
	if err != nil {
		return nil, fmt.Errorf("failed to convert model to turn: %w", err)
	}

	return turn, nil
}

func (r *turnRepository) GetTurnByID(id uuid.UUID, gameID uuid.UUID) (*Turn, error) {
	var model TurnModel

	err := r.db.Where("id = ? AND game_id = ?", id, gameID).First(&model).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("turn not found: id=%s, gameID=%s", id, gameID)
		}
		return nil, fmt.Errorf("failed to get turn by ID: %w", err)
	}

	turn, err := modelToTurn(&model)
	if err != nil {
		return nil, fmt.Errorf("failed to convert model to turn: %w", err)
	}

	return turn, nil
}

func (r *turnRepository) GetAllTurns(gameID uuid.UUID) ([]Turn, error) {
	var models []TurnModel

	err := r.db.Where("game_id = ?", gameID).
		Order("number ASC").
		Find(&models).Error

	if err != nil {
		return nil, fmt.Errorf("failed to get all turns: %w", err)
	}

	turns := make([]Turn, len(models))
	for i, model := range models {
		turn, err := modelToTurn(&model)
		if err != nil {
			return nil, fmt.Errorf("failed to convert model to turn at index %d: %w", i, err)
		}
		turns[i] = *turn
	}

	return turns, nil
}
