package resources

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ResourcesModel struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey"`
	TurnID uuid.UUID `gorm:"type:uuid;not null"`
	Stone  int       `gorm:"not null;default:0"`
	Gold   int       `gorm:"not null;default:0"`
	Wood   int       `gorm:"not null;default:0"`
	Food   int       `gorm:"not null;default:0"`
	Worker int       `gorm:"not null;default:0"`
	gorm.Model
}

func (Resources) TableName() string {
	return "resource"
}

func (g *ResourcesModel) BeforeCreate(tx *gorm.DB) error {
	if g.ID == uuid.Nil {
		g.ID = uuid.New()
	}
	return nil
}

func ResourcesToModel(resources *Resources, turnID uuid.UUID) *ResourcesModel {
	return &ResourcesModel{
		TurnID: turnID,
		Stone:  resources.GetStone(),
		Gold:   resources.GetGold(),
		Wood:   resources.GetWood(),
		Food:   resources.GetFood(),
		Worker: resources.GetWorker(),
	}
}

func ModelToResources(model *ResourcesModel) *Resources {
	return NewResourceControl(
		model.Stone,
		model.Gold,
		model.Wood,
		model.Food,
		model.Worker,
	)
}
