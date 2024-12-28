package usecase

import (
	"fmt"
	"villageQuest/internal/domain/entity/game"
	"villageQuest/internal/infra/repository"
)

type GameStarterUseCase struct {
	gameRepo repository.GameRepository
}

type GameStarter interface {
	Create() (*game.Game, error)
	Load() (*game.Game, error)
}

func NewGameStarter(gameRepo repository.GameRepository) *GameStarterUseCase {
	return &GameStarterUseCase{
		gameRepo: gameRepo,
	}
}

func (s *GameStarterUseCase) getPlayerName() string {
	var playerName string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&playerName)
	return playerName
}

func (s *GameStarterUseCase) loadGamesList() {
	fmt.Println("Games List")
}

func (s *GameStarterUseCase) Create() (game.Game, error) {
	fmt.Println("=== New Game ===")
	playerName := s.getPlayerName()

	nextGameNumber, err := s.gameRepo.GetNextGameNumber()

	gameInstance := game.NewGame(nextGameNumber, playerName)

	s.gameRepo.Insert(*gameInstance)

	return *gameInstance, err
}

func (s *GameStarterUseCase) Load() (game.Game, error) {
	fmt.Println("=== Load Game ===")
	s.loadGamesList()

	return game.Game{}, nil
}
