package game

import (
	"github.com/google/uuid"
)

type Game struct {
	id          uuid.UUID
	number      int
	daysPlayed  int
	playersName string
}

func NewGame(num int, name string) *Game {
	return &Game{
		id:          uuid.New(),
		number:      num,
		daysPlayed:  0,
		playersName: name,
	}
}

func LoadGame(existentId uuid.UUID, num int, days int, name string) *Game {
	return &Game{
		id:          existentId,
		number:      num,
		daysPlayed:  days,
		playersName: name,
	}
}

func (g *Game) Id() uuid.UUID {
	return g.id
}

func (g *Game) Number() int {
	return g.number
}

func (g *Game) DaysPlayed() int {
	return g.daysPlayed
}

func (g *Game) PlayersName() string {
	return g.playersName
}
