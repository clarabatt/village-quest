package game

import (
	"github.com/google/uuid"
)

type Game struct {
	id          uuid.UUID
	number      int
	turnsPlayed int
	playersName string
	isOver      bool
}

func NewGame(num int, name string) *Game {
	return &Game{
		id:          uuid.New(),
		number:      num,
		turnsPlayed: 0,
		playersName: name,
	}
}

func LoadGame(existentId uuid.UUID, num int, turns int, name string) *Game {
	return &Game{
		id:          existentId,
		number:      num,
		turnsPlayed: turns,
		playersName: name,
	}
}

func (g *Game) Id() uuid.UUID {
	return g.id
}

func (g *Game) Number() int {
	return g.number
}

func (g *Game) TurnsPlayed() int {
	return g.turnsPlayed
}

func (g *Game) PlayersName() string {
	return g.playersName
}

func (g *Game) SetOver(isOver bool) {
	g.isOver = isOver
}

func (g *Game) IsOver() bool {
	return g.isOver
}
