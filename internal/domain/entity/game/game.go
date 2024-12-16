package game

import (
	"github.com/google/uuid"
)

type Game struct {
	id            uuid.UUID
	number        int
	maxDaysPlayed int
	playersName   string
}

func NewGame(num int, name string) *Game {
	return &Game{
		id:            uuid.New(),
		number:        num,
		maxDaysPlayed: 0,
		playersName:   name,
	}
}
