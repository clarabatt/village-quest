package usecase

import (
	"fmt"
	"villageQuest/internal/domain/entity/game"
	"villageQuest/internal/infra/repository"
)

type StartNewGame struct {
	gameRepo repository.GameRepository
}

func NewStartNewGame(gameRepo repository.GameRepository) *StartNewGame {
	return &StartNewGame{
		gameRepo: gameRepo,
	}
}

func (s *StartNewGame) CmdInterface() string {
	var playersName string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&playersName)
	return playersName
}

func (s *StartNewGame) Execute() (game.Game, error) {
	player := s.CmdInterface()
	nextGameNumber, err := s.gameRepo.GetNextGameNumber()

	newGame := game.NewGame(nextGameNumber, player)

	s.gameRepo.Insert(*newGame)

	return *newGame, err
}
