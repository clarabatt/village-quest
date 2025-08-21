package menu

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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
		menu.WaitForEnter()
		return
	}

	gameInstance, err := m.GameService.CreateNewGame(playerName)
	if err != nil {
		log.Printf("Error creating game: %v", err)
		fmt.Printf("Failed to create game: %v\n", err)
		menu.WaitForEnter()
		return
	}

	fmt.Printf("Game created successfully for %s!\n", gameInstance.PlayersName())
	fmt.Println("Starting game...")
	menu.WaitForEnter()

	gameLoop := gameplay.NewGameLoop(gameInstance)
	gameLoop.Run()
}

func (m *MainMenu) LoadGame() {
	games, err := m.GameService.GetAllGames()
	if err != nil {
		log.Printf("Error loading games: %v", err)
		fmt.Printf("Failed to load games: %v\n", err)
		menu.WaitForEnter()
		return
	}

	if len(games) == 0 {
		fmt.Println("No saved games found.")
		menu.WaitForEnter()
		return
	}

	mainMenuRef := menu.NewMenu("Main menu", nil)
	loadGameMenu := menu.NewSubMenu("Load Saved Game", mainMenuRef, nil)

	for i, g := range games {
		game := g
		loadGameMenu.AddItem(
			fmt.Sprintf("Game #%d - %s", game.Number(), game.PlayersName()),
			func() {
				fmt.Printf("Loading game #%d for %s...\n", game.Number(), game.PlayersName())
				menu.WaitForEnter()

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
		menu.WaitForEnter()
		return
	}

	if len(games) == 0 {
		fmt.Println("No saved games to delete.")
		menu.WaitForEnter()
		return
	}

	mainMenuRef := menu.NewMenu("Main menu", nil)
	deleteGameMenu := menu.NewSubMenu("Delete Game", mainMenuRef, nil)

	for i, g := range games {
		game := g
		deleteGameMenu.AddItem(
			fmt.Sprintf("Delete Game #%d - %s", game.Number(), game.PlayersName()),
			func() {
				prompt := fmt.Sprintf("Are you sure you want to delete game for %s? (y/n): ", game.PlayersName())

				if menu.GetConfirmation(prompt) {
					if err := m.GameService.DeleteGame(game.Id()); err != nil {
						fmt.Printf("Failed to delete game: %v\n", err)
					} else {
						fmt.Printf("Game for %s deleted successfully!\n", game.PlayersName())
					}
				} else {
					fmt.Println("Delete cancelled.")
				}
				menu.WaitForEnter()
			},
			i+1,
		)
	}

	deleteGameMenu.Show()
}

func getPlayerName() string {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter your name: ")

		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				log.Println("Error reading input:", err)
			}
			return ""
		}

		playerName := strings.TrimSpace(scanner.Text())

		if playerName == "" {
			fmt.Println("Name cannot be empty. Please try again.")
			continue
		}

		if len(playerName) > 50 {
			fmt.Println("Name too long (max 50 characters). Please try again.")
			continue
		}

		if strings.ContainsAny(playerName, "\n\r\t") {
			fmt.Println("Name cannot contain special characters. Please try again.")
			continue
		}

		return playerName
	}
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
	menu.WaitForEnter()
}
