package main

import (
	"villagequest/internal/database"
	"villagequest/internal/domain/game"
	"villagequest/internal/domain/turn"
	"villagequest/internal/engine"
	"villagequest/internal/ui/menu"
)

func main() {
	gormDB := database.NewGormDB()
	defer gormDB.Close()

	gameRepo := game.NewGameRepository(gormDB.DB)
	gameService := game.NewGameService(gameRepo)

	turnRepo := turn.NewTurnRepository(gormDB.DB)
	turnService := turn.NewTurnService(turnRepo)

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
