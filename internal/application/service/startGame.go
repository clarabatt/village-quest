package service

import (
	"fmt"
	"villageQuest/internal/domain/entity/game"
	"villageQuest/internal/infra/repository"
)

type GameStarterService struct {
	gameRepo repository.GameRepository
}

type GameStarter interface {
	Create() (*game.Game, error)
	Load() (*game.Game, error)
}

func NewGameStarterService(gameRepo repository.GameRepository) *GameStarterService {
	return &GameStarterService{
		gameRepo: gameRepo,
	}
}

func (s *GameStarterService) getPlayerName() string {
	var playerName string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&playerName)
	return playerName
}

func (s *GameStarterService) loadGamesList() {
	fmt.Println("Games List")
}

func (s *GameStarterService) Create() (game.Game, error) {
	fmt.Println("=== New Game ===")
	playerName := s.getPlayerName()

	nextGameNumber, err := s.gameRepo.GetNextGameNumber()

	gameInstance := game.NewGame(nextGameNumber, playerName)

	s.gameRepo.Insert(*gameInstance)

	return *gameInstance, err
}

func (s *GameStarterService) Load() (game.Game, error) {
	fmt.Println("=== Load Game ===")
	s.loadGamesList()

	return game.Game{}, nil
}
