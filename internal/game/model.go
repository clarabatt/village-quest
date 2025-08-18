package game

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GameModel struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	Number      int       `gorm:"not null"`
	TurnsPlayed int       `gorm:"default:0"`
	PlayerName  string    `gorm:"type:varchar(255);not null"`
	IsOver      bool      `gorm:"default:false"`
	gorm.Model
}

func (GameModel) TableName() string {
	return "game"
}

func (g *GameModel) BeforeCreate(tx *gorm.DB) error {
	if g.ID == uuid.Nil {
		g.ID = uuid.New()
	}
	return nil
}

func gameToModel(game *Game) *GameModel {
	return &GameModel{
		ID:          game.id,
		Number:      game.number,
		TurnsPlayed: game.turnsPlayed,
		PlayerName:  game.playerName,
		IsOver:      game.isOver,
	}
}

func modelToGame(model *GameModel) *Game {
	return &Game{
		id:          model.ID,
		number:      model.Number,
		turnsPlayed: model.TurnsPlayed,
		playerName:  model.PlayerName,
		isOver:      model.IsOver,
	}
}
