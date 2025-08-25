package engine

import (
	"fmt"
	"strconv"
	"villagequest/internal/game"
	"villagequest/internal/ui/menu"
)

type GameLoop struct {
	Game *game.Game
	Turn int
}

type GameRunner interface {
    Run()
}

func NewGameLoop(g *game.Game) *GameLoop {
	return &GameLoop{Game: g, Turn: 1}
}

func (loop *GameLoop) Run() {
	for {
		m := menu.NewMenu("Turn "+strconv.Itoa(loop.Turn), nil)
		m.AddItem("Build", loop.Build, 1)
		m.AddItem("Quit Game", loop.Quit, 4)
		m.Show()
		loop.Turn += 1

		if loop.Game.IsOver() {
			break
		}
	}
}

func (g *GameLoop) Build() {
	fmt.Println("You built a house")
}

func (g *GameLoop) Quit() {
	fmt.Println("Thanks for playing!")
	g.Game.SetOver(true)
}
