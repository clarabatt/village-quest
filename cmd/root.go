package cmd

import (
	"fmt"
	"villageQuest/internal/application/usecase"
	"villageQuest/internal/infra/database"
	"villageQuest/internal/infra/repository"
)

func Execute() {
	fmt.Println("=== Village Quest ===\n")
	fmt.Println("Start a new game: ")

	dbConnection := database.NewSqliteAdapter()
	gameRepo := repository.NewGameRepository(dbConnection)

	startNewGame := usecase.NewStartNewGame(gameRepo)
	startNewGame.Execute()
}
