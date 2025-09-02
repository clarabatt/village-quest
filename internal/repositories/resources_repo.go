package repositories

import (
	"errors"
	"fmt"

	. "villagequest/internal/domain/resources"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ResourcesRepository interface {
	Create(resources *Resources, turnID uuid.UUID) (*Resources, error)
	GetByID(id uuid.UUID) (*Resources, error)
	GetByTurnId(id uuid.UUID) (*Resources, error)
}

type resourcesRepository struct {
	db *gorm.DB
}

func NewResourcesRepository(db *gorm.DB) ResourcesRepository {
	return &resourcesRepository{
		db: db,
	}
}

func (r *resourcesRepository) Create(resources *Resources, turnID uuid.UUID) (*Resources, error) {
	model := ResourcesToModel(resources, turnID)
	if err := r.db.Create(model).Error; err != nil {
		return nil, fmt.Errorf("failed to create resources: %w", err)
	}
	result := ModelToResources(model)
	return result, nil
}

func (r *resourcesRepository) GetByID(id uuid.UUID) (*Resources, error) {
	var model ResourcesModel

	err := r.db.Where("id = ?", id).First(&model).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("resources not found: id=%s", id)
		}
		return nil, fmt.Errorf("failed to get resources by ID: %w", err)
	}

	game := ModelToResources(&model)
	return game, nil
}

func (r *resourcesRepository) GetByTurnId(id uuid.UUID) (*Resources, error) {
	var model ResourcesModel

	err := r.db.Where("turn_id = ?", id).First(&model).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("resources not found: turn_id=%s", id)
		}
		return nil, fmt.Errorf("failed to get resources by turn ID: %w", err)
	}

	game := ModelToResources(&model)
	return game, nil
}
