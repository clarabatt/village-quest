package turn

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TurnModel struct {
	ID                 uuid.UUID `gorm:"type:uuid;primaryKey"`
	GameID             uuid.UUID `gorm:"type:uuid;index;not null"`
	Number             int       `gorm:"not null"`
	Status             string    `gorm:"type:varchar(20);not null;default:'in_progress'"`
	CurrentPhase       string    `gorm:"type:varchar(20);not null;default:'collect'"`
	ActionsUsed        int       `gorm:"not null;default:0"`
	ResourcesCollected bool      `gorm:"not null;default:false"`
	EventsProcessed    bool      `gorm:"not null;default:false"`
	gorm.Model
}

func (t *TurnModel) BeforeCreate(tx *gorm.DB) error {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}
	return nil
}

func TurnToModel(turn *Turn) (*TurnModel, error) {
	return &TurnModel{
		ID:                 turn.GetID(),
		GameID:             turn.GetGameID(),
		Number:             turn.GetNumber(),
		Status:             string(turn.GetStatus()),
		CurrentPhase:       string(turn.GetCurrentPhase()),
		ActionsUsed:        turn.GetActionsUsed(),
		ResourcesCollected: turn.AreResourcesCollected(),
		EventsProcessed:    turn.AreEventsProcessed(),
	}, nil
}

func ModelToTurn(model *TurnModel) (*Turn, error) {
	return LoadTurn(
		model.ID,
		model.GameID,
		model.Number,
		TurnStatus(model.Status),
		TurnPhase(model.CurrentPhase),
		model.ActionsUsed,
		model.ResourcesCollected,
		model.EventsProcessed,
	), nil
}

func (TurnModel) TableName() string {
	return "turn"
}
