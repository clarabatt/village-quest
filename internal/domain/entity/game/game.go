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

func NewSavedGame(existentId uuid.UUID, num int, days int, name string) *Game {
	return &Game{
		id:            existentId,
		number:        num,
		maxDaysPlayed: days,
		playersName:   name,
	}
}

func (g *Game) Id() uuid.UUID {
	return g.id
}

func (g *Game) Number() int {
	return g.number
}

func (g *Game) MaxDaysPlayed() int {
	return g.maxDaysPlayed
}

func (g *Game) PlayersName() string {
	return g.playersName
}
