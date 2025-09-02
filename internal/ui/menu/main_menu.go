package menu

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	. "villagequest/internal/application"
	"villagequest/internal/domain/game"
)

type MainMenu struct {
	GameService GameService
	GameStarter func(*game.Game)
}

func RunMainMenu(gameService GameService, gameStarter func(*game.Game)) {
	DisplayWelcome()

	m := NewMenu("Main menu", nil)
	mainMenu := &MainMenu{
		GameService: gameService,
		GameStarter: gameStarter,
	}

	m.AddItem("New game", mainMenu.NewGame, 1)
	m.AddItem("Load game", mainMenu.LoadGame, 2)
	m.AddItem("Delete game", mainMenu.DeleteGame, 3)

	m.Show()
}

func (m *MainMenu) startGame(gameInstance *game.Game) {
	if m.GameStarter != nil {
		m.GameStarter(gameInstance)
	}
}

func (m *MainMenu) NewGame() {
	playerName := getPlayerName()
	if playerName == "" {
		fmt.Println("Invalid player name. Please try again.")
		WaitForEnter()
		return
	}

	gameInstance, err := m.GameService.CreateNewGame(playerName)
	if err != nil {
		log.Printf("Error creating game: %v", err)
		fmt.Printf("Failed to create game: %v\n", err)
		WaitForEnter()
		return
	}

	fmt.Printf("Game created successfully for %s!\n", gameInstance.PlayersName())
	fmt.Println("Starting game...")
	WaitForEnter()

	m.startGame(gameInstance)
}

func (m *MainMenu) LoadGame() {
	games, err := m.GameService.GetAllGames()
	if err != nil {
		log.Printf("Error loading games: %v", err)
		fmt.Printf("Failed to load games: %v\n", err)
		WaitForEnter()
		return
	}

	if len(games) == 0 {
		fmt.Println("No saved games found.")
		WaitForEnter()
		return
	}

	mainMenuRef := NewMenu("Main menu", nil)
	loadGameMenu := NewSubMenu("Load Saved Game", mainMenuRef, nil)

	for i, g := range games {
		game := g
		loadGameMenu.AddItem(
			fmt.Sprintf("Game #%d - %s", game.Number(), game.PlayersName()),
			func() {
				fmt.Printf("Loading game #%d for %s...\n", game.Number(), game.PlayersName())
				WaitForEnter()

				m.startGame(&game)
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
		WaitForEnter()
		return
	}

	if len(games) == 0 {
		fmt.Println("No saved games to delete.")
		WaitForEnter()
		return
	}

	mainMenuRef := NewMenu("Main menu", nil)
	deleteGameMenu := NewSubMenu("Delete Game", mainMenuRef, nil)

	for i, g := range games {
		game := g
		deleteGameMenu.AddItem(
			fmt.Sprintf("Delete Game #%d - %s", game.Number(), game.PlayersName()),
			func() {
				prompt := fmt.Sprintf("Are you sure you want to delete game for %s? (y/n): ", game.PlayersName())

				if GetConfirmation(prompt) {
					if err := m.GameService.DeleteGame(game.Id()); err != nil {
						fmt.Printf("Failed to delete game: %v\n", err)
					} else {
						fmt.Printf("Game for %s deleted successfully!\n", game.PlayersName())
					}
				} else {
					fmt.Println("Delete cancelled.")
				}
				WaitForEnter()
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
	WaitForEnter()
}
