package application

import (
	. "villagequest/internal/domain/turn"
	. "villagequest/internal/repositories"

	"github.com/google/uuid"
)

type turnService struct {
	turnRepo TurnRepository
}

type TurnService interface {
	LoadLastTurn(gameId uuid.UUID) *Turn
	LoadFirstTurn(gameId uuid.UUID) *Turn
	FinishTurn(turn Turn)
	StartNextTurn(turn Turn) *Turn
}

func NewTurnService(repository TurnRepository) TurnService {
	return &turnService{
		turnRepo: repository,
	}
}

func (ts *turnService) LoadLastTurn(gameId uuid.UUID) *Turn {
	lastTurn, err := ts.turnRepo.GetLastTurn(gameId)
	if err != nil {
		return nil
	}
	return lastTurn
}

func (ts *turnService) LoadFirstTurn(gameId uuid.UUID) *Turn {
	newTurn := CreateTurn(gameId)
	result, err := ts.turnRepo.Create(newTurn)
	if err != nil {
		return nil
	}
	return result
}

func (ts *turnService) FinishTurn(turn Turn) {

}

func (ts *turnService) StartNextTurn(turn Turn) *Turn {
	nextTurn := turn.CreateNextTurn()
	result, err := ts.turnRepo.Create(nextTurn)
	if err != nil {
		return nil
	}
	return result
}
