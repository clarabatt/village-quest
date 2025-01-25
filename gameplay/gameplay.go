package gameplay

import (
	"fmt"
	"villageQuest/internal/application/usecase"
	"villageQuest/internal/infra/database"
	"villageQuest/internal/infra/repository"
)

func Execute() {
	dbConnection := database.NewSqliteAdapter()
	gameRepo := repository.NewGameRepository(dbConnection)
	gameStarterService := usecase.NewGameStarter(gameRepo)

	for {
		r := mainMenu(gameStarterService)
		if r == 0 {
			break
		}
	}

}

func mainMenu(starterService *usecase.GameStarterUseCase) int {
	fmt.Println("=== Welcome to Village Quest ===")
	fmt.Println("1. Start a new game")
	fmt.Println("2. Load a game")
	fmt.Println("3. Exit")

	fmt.Print("> ")
	var option int
	var selectedGame int
	fmt.Scanln(&option)

	switch option {
	case 1:
		fmt.Println("=== New Game ===")
		playerName := getPlayerName()
		starterService.Create(playerName)
		return 1
	case 2:
		fmt.Println("=== Load Game ===")
		games, _ := starterService.GetAllGames()
		for _, game := range games {
			fmt.Printf("%d. %s\n", game.Number(), game.PlayersName())
		}
		fmt.Print("> ")
		fmt.Scanln(&selectedGame)
		fmt.Println("Loading game... ", games[selectedGame-1].PlayersName())
		return 1
	case 3:
		fmt.Println("Goodbye!")
		return 0
	default:
		fmt.Println("Invalid option. Try again.")
		return 1
	}
}

func getPlayerName() string {
	var playerName string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&playerName)
	return playerName
}
