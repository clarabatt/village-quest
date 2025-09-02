package engine

import (
	"fmt"
	"strings"
	. "villagequest/internal/application"
	. "villagequest/internal/domain/game"

	"villagequest/internal/ui/menu"
)

type GameLoop struct {
	Game        *Game
	Turn        int
	GameService GameService
	TurnService TurnService
}

type GameRunner interface {
	Run()
}

func NewGameLoop(g *Game, gameService GameService, turnService TurnService) *GameLoop {
	return &GameLoop{
		Game:        g,
		Turn:        1,
		GameService: gameService,
		TurnService: turnService,
	}
}

func (loop *GameLoop) Run() {
	for {
		loop.displayGameStatus()
		menuTitle := fmt.Sprintf("🏘️ Village Quest - Year %d 🏘️", loop.Turn)
		m := menu.NewMenu(menuTitle, nil)

		m.AddItem("🏠 Build Structure", loop.Build, 1)
		m.AddItem("❌ Quit Game", loop.Quit, 9)

		m.Show()

		// TODO: Increment turn

		if err := loop.GameService.SaveGame(loop.Game); err != nil {
			fmt.Printf("Warning: Failed to auto-save game: %v\n", err)
		}

		if loop.Game.IsOver() {
			loop.displayGameOver()
			break
		}
	}
}

func (loop *GameLoop) displayGameOver() {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("🎉 GAME OVER 🎉")
	fmt.Printf("Your village lasted %d years!\n", loop.Turn-1)
	fmt.Printf("Thanks for playing, %s!\n", loop.Game.PlayersName())
	fmt.Println(strings.Repeat("=", 60))
}

func (loop *GameLoop) displayGameStatus() {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Printf("Player: %s | Year: %d\n", loop.Game.PlayersName(), loop.Turn)
	fmt.Println(strings.Repeat("=", 60))
}

func (loop *GameLoop) Build() {
	fmt.Println("You built a house")
}

func (loop *GameLoop) Quit() {
	fmt.Println("Thanks for playing!")
	fmt.Println("💾 Saving your progress...")

	if err := loop.GameService.SaveGame(loop.Game); err != nil {
		fmt.Printf("❌ Error saving game: %v\n", err)
	} else {
		fmt.Println("✅ Game saved successfully!")
	}

	loop.Game.SetOver(true)
}
