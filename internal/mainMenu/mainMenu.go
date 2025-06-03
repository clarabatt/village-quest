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
	GameService game.GameService
}

func Execute() {
	gormDB := database.NewGormDB()
	defer gormDB.Close()

	gameRepo := game.NewGameRepository(gormDB.DB)
	gameService := game.NewGameService(gameRepo)

	// turnRepo := turn.NewTurnRepository(gormDB)
	// turnService := turn.NewTurnService(turnRepo)

	RunMainMenu(gameService)
}

func RunMainMenu(gameService game.GameService) {
	DisplayWelcome()
	fmt.Scanln()

	m := menu.NewMenu("Main menu", nil)
	mainMenu := &MainMenu{GameService: gameService}

	m.AddItem("New game", mainMenu.NewGame, 1)
	m.AddItem("Load game", mainMenu.LoadGame, 2)
	m.AddItem("Delete game", mainMenu.DeleteGame, 3)

	m.Show()
}

func (m *MainMenu) NewGame() {
	playerName := getPlayerName()
	if playerName == "" {
		fmt.Println("Invalid player name. Please try again.")
		fmt.Scanln()
		return
	}

	gameInstance, err := m.GameService.CreateNewGame(playerName)
	if err != nil {
		log.Printf("Error creating game: %v", err)
		fmt.Printf("Failed to create game: %v\n", err)
		fmt.Scanln()
		return
	}

	fmt.Printf("Game created successfully for %s!\n", gameInstance.PlayersName())
	fmt.Println("Starting game...")
	fmt.Scanln()

	gameLoop := gameplay.NewGameLoop(gameInstance)
	gameLoop.Run()
}

func (m *MainMenu) LoadGame() {
	games, err := m.GameService.GetAllGames()
	if err != nil {
		log.Printf("Error loading games: %v", err)
		fmt.Printf("Failed to load games: %v\n", err)
		fmt.Scanln()
		return
	}

	if len(games) == 0 {
		fmt.Println("No saved games found.")
		fmt.Println("Press Enter to continue...")
		fmt.Scanln()
		return
	}

	loadGameMenu := menu.NewMenu("Load Saved Game", nil)

	for i, g := range games {
		game := g
		loadGameMenu.AddItem(
			fmt.Sprintf("Game #%d - %s", game.Number(), game.PlayersName()),
			func() {
				fmt.Printf("Loading game #%d for %s...\n", game.Number(), game.PlayersName())
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

func (m *MainMenu) DeleteGame() {
	games, err := m.GameService.GetAllGames()
	if err != nil {
		log.Printf("Error loading games: %v", err)
		fmt.Printf("Failed to load games: %v\n", err)
		fmt.Scanln()
		return
	}

	if len(games) == 0 {
		fmt.Println("No saved games to delete.")
		fmt.Println("Press Enter to continue...")
		fmt.Scanln()
		return
	}

	deleteGameMenu := menu.NewMenu("Delete Game", nil)

	for i, g := range games {
		game := g
		deleteGameMenu.AddItem(
			fmt.Sprintf("Delete Game #%d - %s", game.Number(), game.PlayersName()),
			func() {
				fmt.Printf("Are you sure you want to delete game for %s? (y/n): ", game.PlayersName())
				var response string
				fmt.Scanln(&response)

				if response == "y" || response == "Y" {
					if err := m.GameService.DeleteGame(game.Id()); err != nil {
						fmt.Printf("Failed to delete game: %v\n", err)
					} else {
						fmt.Printf("Game for %s deleted successfully!\n", game.PlayersName())
					}
				} else {
					fmt.Println("Delete cancelled.")
				}
				fmt.Println("Press Enter to continue...")
				fmt.Scanln()
			},
			i+1,
		)
	}

	deleteGameMenu.Show()
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
