package game

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GameRepository interface {
	Create(game *Game) (*Game, error)
	GetByID(id uuid.UUID) (*Game, error)
	GetByNumber(number int) (*Game, error)
	Update(game *Game) (*Game, error)
	Delete(id uuid.UUID) error
	GetAll() ([]Game, error)
}

type gameRepository struct {
	db *gorm.DB
}

func NewGameRepository(db *gorm.DB) GameRepository {
	return &gameRepository{
		db: db,
	}
}

func (r *gameRepository) Create(game *Game) (*Game, error) {
	model := gameToModel(game)

	if err := r.db.Create(model).Error; err != nil {
		return nil, fmt.Errorf("failed to create game: %w", err)
	}

	result := modelToGame(model)
	return result, nil
}

func (r *gameRepository) GetByID(id uuid.UUID) (*Game, error) {
	var model GameModel

	err := r.db.Where("id = ?", id).First(&model).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("game not found: id=%s", id)
		}
		return nil, fmt.Errorf("failed to get game by ID: %w", err)
	}

	game := modelToGame(&model)
	return game, nil
}

func (r *gameRepository) GetByNumber(number int) (*Game, error) {
	var model GameModel

	err := r.db.Where("number = ?", number).First(&model).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("game not found: number=%d", number)
		}
		return nil, fmt.Errorf("failed to get game by number: %w", err)
	}

	game := modelToGame(&model)
	return game, nil
}

func (r *gameRepository) Update(game *Game) (*Game, error) {
	model := gameToModel(game)

	result := r.db.Model(&GameModel{}).
		Where("id = ?", game.id).
		Updates(model)

	if result.Error != nil {
		return nil, fmt.Errorf("failed to update game: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("game not found: id=%s", game.id)
	}

	var updatedModel GameModel
	if err := r.db.Where("id = ?", game.id).First(&updatedModel).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch updated game: %w", err)
	}

	updatedGame := modelToGame(&updatedModel)
	return updatedGame, nil
}

func (r *gameRepository) Delete(id uuid.UUID) error {
	result := r.db.Delete(&GameModel{}, "id = ?", id)

	if result.Error != nil {
		return fmt.Errorf("failed to delete game: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("game not found: id=%s", id)
	}

	return nil
}

func (r *gameRepository) GetAll() ([]Game, error) {
	var models []GameModel

	err := r.db.Order("created_at ASC").Find(&models).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get all games: %w", err)
	}

	games := make([]Game, len(models))
	for i, model := range models {
		games[i] = *modelToGame(&model)
	}

	return games, nil
}
