package main

import (
	. "villagequest/internal/application"
	"villagequest/internal/database"
	"villagequest/internal/domain/game"
	"villagequest/internal/engine"
	. "villagequest/internal/repositories"
	"villagequest/internal/ui/menu"
)

func main() {
	gormDB := database.NewGormDB()
	defer gormDB.Close()

	gameRepo := NewGameRepository(gormDB.DB)
	gameService := NewGameService(gameRepo)

	turnRepo := NewTurnRepository(gormDB.DB)
	turnService := NewTurnService(turnRepo)

	gameStarter := func(g *game.Game) {
		gameLoop := engine.NewGameLoop(
			g,
			gameService,
			turnService,
		)
		gameLoop.Run()
	}

	menu.RunMainMenu(gameService, gameStarter)
}
