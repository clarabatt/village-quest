package mainMenu

import (
	"fmt"
	"log"
	"villagequest/internal/database"
	"villagequest/internal/game"
	"villagequest/internal/menu"
)

type MainMenu struct {
	StarterService *game.GameStarterUseCase
}

func Execute() {
	dbConnection := database.NewSqliteAdapter()
	gameRepo := game.NewGameRepository(dbConnection)
	gameStarterService := game.NewGameStarter(gameRepo)

	RunMainMenu(gameStarterService)
}

func RunMainMenu(starterService *game.GameStarterUseCase) {
	m := menu.NewMenu("Welcome to Village Quest!", nil)
	mainMenu := &MainMenu{StarterService: starterService}
	m.AddItem("New game", mainMenu.NewGame, 1)
	m.AddItem("Load game", mainMenu.LoadGame, 2)

	m.Show()
}

func (m *MainMenu) NewGame() {
	playerName := getPlayerName()
	if _, err := m.StarterService.Create(playerName); err != nil {
		log.Println("Error creating game:", err)
	}
}

func (m *MainMenu) LoadGame() {
	games, err := m.StarterService.GetAllGames()
	if err != nil {
		log.Println("Error loading games:", err)
		return
	}

	if len(games) == 0 {
		fmt.Println("No saved games.")
		fmt.Scanln()
		return
	}

	loadGameMenu := menu.NewMenu("Load Saved Game", nil)

	for i, g := range games {
		game := g
		_ = loadGameMenu.AddItem(
			fmt.Sprintf("%s", game.PlayersName()),
			func() {
				fmt.Printf("Loading game for %s...\n", game.PlayersName())
				fmt.Println("Press Enter to continue...")
				fmt.Scanln()
				// TODO: Pass game to the next phase (e.g., game loop)
			},
			i+1,
		)
	}
	loadGameMenu.Show()
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
