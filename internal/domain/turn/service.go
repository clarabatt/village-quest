package turn

import "github.com/google/uuid"

type turnService struct {
	repository TurnRepository
}

type TurnService interface {
	LoadLastTurn(gameId uuid.UUID) *Turn
	LoadFirstTurn(gameId uuid.UUID) *Turn
	FinishTurn(turn Turn)
	StartNextTurn(turn Turn) *Turn
}

func NewTurnService(repository TurnRepository) TurnService {
	return &turnService{
		repository: repository,
	}
}

func (ts *turnService) LoadLastTurn(gameId uuid.UUID) *Turn {
	lastTurn, err := ts.repository.GetLastTurn(gameId)
	if err != nil {
		return nil
	}
	return lastTurn
}

func (ts *turnService) LoadFirstTurn(gameId uuid.UUID) *Turn {
	newTurn := CreateTurn(gameId)
	result, err := ts.repository.Create(newTurn)
	if err != nil {
		return nil
	}
	return result
}

func (ts *turnService) FinishTurn(turn Turn) {

}

func (ts *turnService) StartNextTurn(turn Turn) *Turn {
	nextTurn := turn.CreateNextTurn()
	result, err := ts.repository.Create(nextTurn)
	if err != nil {
		return nil
	}
	return result
}
