package turn

import (
	"encoding/json"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TurnModel struct {
	ID               uuid.UUID `gorm:"type:uuid;primaryKey"`
	GameID           uuid.UUID `gorm:"type:uuid;index;not null"`
	Number           int       `gorm:"not null"`
	PlayersAction    *string   `gorm:"type:varchar(50)"`
	RandomEvents     string    `gorm:"type:text"` // JSON as string
	InitialResources string    `gorm:"type:text"` // JSON as string
	FinalResources   string    `gorm:"type:text"` // JSON as string
	gorm.Model
}

func (t *TurnModel) BeforeCreate(tx *gorm.DB) error {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}
	return nil
}

func turnToModel(turn *Turn, gameID uuid.UUID) (*TurnModel, error) {
	eventsJSON, err := json.Marshal(turn.RandomEvents)
	if err != nil {
		return nil, err
	}

	initialResourcesJSON, err := json.Marshal(turn.InitialResources)
	if err != nil {
		return nil, err
	}

	finalResourcesJSON, err := json.Marshal(turn.FinalResources)
	if err != nil {
		return nil, err
	}

	var actionStr *string
	if turn.PlayersAction != nil {
		s := string(*turn.PlayersAction)
		actionStr = &s
	}

	return &TurnModel{
		ID:               turn.Id,
		GameID:           gameID,
		Number:           turn.Number,
		PlayersAction:    actionStr,
		RandomEvents:     string(eventsJSON),
		InitialResources: string(initialResourcesJSON),
		FinalResources:   string(finalResourcesJSON),
	}, nil
}

func modelToTurn(model *TurnModel) (*Turn, error) {
	var events []Event
	if err := json.Unmarshal([]byte(model.RandomEvents), &events); err != nil {
		return nil, err
	}

	var initialResources []string
	if err := json.Unmarshal([]byte(model.InitialResources), &initialResources); err != nil {
		return nil, err
	}

	var finalResources []string
	if err := json.Unmarshal([]byte(model.FinalResources), &finalResources); err != nil {
		return nil, err
	}

	var action *Action
	if model.PlayersAction != nil {
		a := Action(*model.PlayersAction)
		action = &a
	}

	return &Turn{
		Id:               model.ID,
		Number:           model.Number,
		PlayersAction:    action,
		RandomEvents:     events,
		InitialResources: initialResources,
		FinalResources:   finalResources,
	}, nil
}
