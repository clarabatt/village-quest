package mainMenu

import (
	"fmt"
	"log"
	"villagequest/internal/database"
	"villagequest/internal/game"
	"villagequest/internal/gameplay"
	"villagequest/internal/menu"
)

type MainMenu struct {
	StarterService *game.GameStarterUseCase
}

func Execute() {
	dbConnection := database.NewSqliteAdapter()
	gameRepo := game.NewGameRepository(dbConnection)
	gameStarterService := game.NewGameStarter(gameRepo)

	RunMainMenu(gameStarterService)
}

func RunMainMenu(starterService *game.GameStarterUseCase) {
	DisplayWelcome()
	fmt.Scanln()
	m := menu.NewMenu("Main menu", nil)
	mainMenu := &MainMenu{StarterService: starterService}
	m.AddItem("New game", mainMenu.NewGame, 1)
	m.AddItem("Load game", mainMenu.LoadGame, 2)

	m.Show()
}

func (m *MainMenu) NewGame() {
	playerName := getPlayerName()
	if _, err := m.StarterService.Create(playerName); err != nil {
		log.Println("Error creating game:", err)
	}
}

func (m *MainMenu) LoadGame() {
	games, err := m.StarterService.GetAllGames()
	if err != nil {
		log.Println("Error loading games:", err)
		return
	}

	if len(games) == 0 {
		fmt.Println("No saved games.")
		fmt.Scanln()
		return
	}

	loadGameMenu := menu.NewMenu("Load Saved Game", nil)

	for i, g := range games {
		game := g
		_ = loadGameMenu.AddItem(
			fmt.Sprintf("%s", game.PlayersName()),
			func() {
				fmt.Printf("Loading game for %s...\n", game.PlayersName())
				fmt.Println("Press Enter to continue...")
				fmt.Scanln()

				gameLoop := gameplay.NewGameLoop(&game)
				gameLoop.Run()
			},
			i+1,
		)
	}
	loadGameMenu.Show()
}

func getPlayerName() string {
	var playerName string
	fmt.Print("Enter your name: ")
	if _, err := fmt.Scanln(&playerName); err != nil {
		log.Println("Error reading input:", err)
		return ""
	}
	return playerName
}

func DisplayWelcome() {
	fmt.Println("╔═══════════════════════════════════════════════════════════════╗")
	fmt.Println("║                                                               ║")
	fmt.Println("║                     🏘️  VILLAGE QUEST 🏘️                        ║")
	fmt.Println("║                                                               ║")
	fmt.Println("║              Build your settlement from scratch!              ║")
	fmt.Println("║                                                               ║")
	fmt.Println("║  🎯 Goal: Transform your small village into a thriving city   ║")
	fmt.Println("║  ⏰ Each turn = 1 year, 2 actions per turn                    ║")
	fmt.Println("║  📊 Manage resources, build structures, handle events         ║")
	fmt.Println("║                                                               ║")
	fmt.Println("╚═══════════════════════════════════════════════════════════════╝")
	fmt.Print("\nPress Enter to start your village...")
}
