package gameplay

import (
	"fmt"
	"log"
	"villageQuest/internal/database"
	"villageQuest/internal/game"
)

func Execute() {
	dbConnection := database.NewSqliteAdapter()
	gameRepo := game.NewGameRepository(dbConnection)
	gameStarterService := game.NewGameStarter(gameRepo)
	for {
		r := mainMenu(gameStarterService)
		if r == 0 {
			break
		}
	}
}

func mainMenu(starterService *game.GameStarterUseCase) int {
	fmt.Println("=== Welcome to Village Quest ===")
	fmt.Println("1. Start a new game")
	fmt.Println("2. Load a game")
	fmt.Println("0. Exit")

	fmt.Print("> ")
	var option int
	var selectedGame int
	if _, err := fmt.Scanln(&option); err != nil {
		log.Println("Error reading input:", err)
		return 1
	}

	switch option {
	case 0:
		fmt.Println("Goodbye!")
		return 0
	case 1:
		fmt.Println("=== New Game ===")
		playerName := getPlayerName()
		if _, err := starterService.Create(playerName); err != nil {
			log.Println("Error creating game:", err)
		}
		return 1
	case 2:
		fmt.Println("=== Load Game ===")
		games, _ := starterService.GetAllGames()
		for _, game := range games {
			fmt.Printf("%d. %s\n", game.Number(), game.PlayersName())
		}
		fmt.Print("> ")
		if _, err := fmt.Scanln(&selectedGame); err != nil {
			log.Println("Error reading input:", err)
		}
		fmt.Println("Loading game... ", games[selectedGame-1].PlayersName())
		return 1
	default:
		fmt.Println("Invalid option. Try again.")
		return 1
	}
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
