package gameplay

import (
	"fmt"
	"villageQuest/internal/application/service"
	"villageQuest/internal/infra/database"
	"villageQuest/internal/infra/repository"
)

func Execute() {
	dbConnection := database.NewSqliteAdapter()
	gameRepo := repository.NewGameRepository(dbConnection)
	gameStarterService := service.NewGameStarterService(gameRepo)

	mainMenu(gameStarterService)
}

func mainMenu(starterService *service.GameStarterService) {
	fmt.Println("=== Village Quest ===")
	fmt.Println("1. Start a new game")
	fmt.Println("2. Load a game")
	fmt.Println("3. Exit")

	fmt.Print("> ")
	var option int
	fmt.Scanln(&option)

	switch option {
	case 1:
		starterService.Create()
	case 2:
		starterService.Load()
	case 3:
		fmt.Println("Goodbye!")
	}
}
