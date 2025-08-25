package main

import (
	"villagequest/internal/database"
	"villagequest/internal/engine"
	"villagequest/internal/game"
	"villagequest/internal/ui/menu"
)

func main() {
	gormDB := database.NewGormDB()
	defer gormDB.Close()

	gameRepo := game.NewGameRepository(gormDB.DB)
	gameService := game.NewGameService(gameRepo)

	// turnRepo := turn.NewTurnRepository(gormDB)
	// turnService := turn.NewTurnService(turnRepo)

	gameStarter := func(g *game.Game) {
        gameLoop := engine.NewGameLoop(g)
        gameLoop.Run()
    }

	menu.RunMainMenu(gameService, gameStarter)
}
