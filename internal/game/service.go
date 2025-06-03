package game

import (
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
)

var (
	ErrGameNotFound    = errors.New("game not found")
	ErrEmptyId         = errors.New("game ID cannot be empty")
	ErrPlayerNameEmpty = errors.New("player name cannot be empty")
)

type gameService struct {
	repository GameRepository
}

type GameService interface {
	CreateNewGame(playerName string) (*Game, error)
	GetAllGames() ([]Game, error)
	DeleteGame(gameId uuid.UUID) error
}

func NewGameService(repository GameRepository) GameService {
	return &gameService{
		repository: repository,
	}
}

func (s *gameService) CreateNewGame(playerName string) (*Game, error) {
	if playerName == "" {
		return nil, ErrPlayerNameEmpty
	}

	existingGames, err := s.repository.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch existing games: %w", err)
	}

	nextGameNumber := len(existingGames) + 1

	gameInstance := NewGame(nextGameNumber, playerName)

	if _, err := s.repository.Create(gameInstance); err != nil {
		return nil, fmt.Errorf("failed to create game: %w", err)
	}

	log.Printf("Created new game for player: %s (Game #%d)", playerName, nextGameNumber)
	return gameInstance, nil
}

func (s *gameService) GetAllGames() ([]Game, error) {
	games, err := s.repository.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch games: %w", err)
	}

	return games, nil
}

func (s *gameService) DeleteGame(gameId uuid.UUID) error {
	_, err := s.repository.GetByID(gameId)
	if err != nil {
		return ErrGameNotFound
	}

	err = s.repository.Delete(gameId)
	if err != nil {
		return fmt.Errorf("failed to delete game: %w", err)
	}

	log.Printf("Deleted game with ID: %s", gameId)
	return nil
}
