package usecase

import (
	"log"
	"villageQuest/domain/entity/game"
	"villageQuest/repository"
)

type GameStarterUseCase struct {
	gameRepo  repository.GameRepository
	gamesList []game.Game
}

type GameStarter interface {
	Create() (*game.Game, error)
	GetAllGames() (*game.Game, error)
}

func NewGameStarter(gameRepo repository.GameRepository) *GameStarterUseCase {
	games, _ := gameRepo.GetAll()
	return &GameStarterUseCase{
		gameRepo:  gameRepo,
		gamesList: games,
	}
}

func (s *GameStarterUseCase) Create(playerName string) (game.Game, error) {
	nextGameNumber, err := s.gameRepo.GetNextGameNumber()
	gameInstance := game.NewGame(nextGameNumber, playerName)
	if err := s.gameRepo.Insert(*gameInstance); err != nil {
		log.Print("Error inserting a game instance")
	}
	s.loadGamesList()
	return *gameInstance, err
}

func (s *GameStarterUseCase) GetAllGames() ([]game.Game, error) {
	return s.gamesList, nil
}

func (s *GameStarterUseCase) loadGamesList() {
	games, _ := s.gameRepo.GetAll()
	s.gamesList = games
}
