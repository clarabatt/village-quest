package game

import (
	"log"
)

type GameStarterUseCase struct {
	gameRepo  GameRepository
	gamesList []Game
}

type GameStarter interface {
	Create() (*Game, error)
	GetAllGames() (*Game, error)
}

func NewGameStarter(gameRepo GameRepository) *GameStarterUseCase {
	games, _ := gameRepo.GetAll()
	return &GameStarterUseCase{
		gameRepo:  gameRepo,
		gamesList: games,
	}
}

func (s *GameStarterUseCase) Create(playerName string) (Game, error) {
	var err error
	nextGameNumber := len(s.gamesList) + 1
	gameInstance := NewGame(nextGameNumber, playerName)
	if err := s.gameRepo.Insert(*gameInstance); err != nil {
		log.Print("Error inserting a game instance")
	}
	s.loadGamesList()
	return *gameInstance, err
}

func (s *GameStarterUseCase) GetAllGames() ([]Game, error) {
	return s.gamesList, nil
}

func (s *GameStarterUseCase) loadGamesList() {
	games, _ := s.gameRepo.GetAll()
	s.gamesList = games
}
