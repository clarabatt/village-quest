package turn

import "github.com/google/uuid"

type turnService struct {
	repository TurnRepository
}

type TurnService interface {
	StartFirstTurn(gameId uuid.UUID) (*Turn, error)

	GetLatestTurn(gameId uuid.UUID) (*Turn, error)
	AdvanceToNextTurn(gameId uuid.UUID) (*Turn, error)
	CompleteTurn(gameId uuid.UUID, action Action, events []Event, finalResources []string) (*Turn, error)

	SetPlayerAction(gameId uuid.UUID, action Action) (*Turn, error)
	AddRandomEvent(gameId uuid.UUID, event Event) (*Turn, error)
	UpdateResources(gameId uuid.UUID, finalResources []string) (*Turn, error)
}

func NewTurnService(repository TurnRepository) TurnService {
	return &turnService{
		repository: repository,
	}
}

func (s *turnService) StartFirstTurn(gameId uuid.UUID) (*Turn, error) {
	return nil, nil
}
func (s *turnService) GetLatestTurn(gameId uuid.UUID) (*Turn, error) {
	return nil, nil
}
func (s *turnService) AdvanceToNextTurn(gameId uuid.UUID) (*Turn, error) {
	return nil, nil
}
func (s *turnService) CompleteTurn(gameId uuid.UUID, action Action, events []Event, finalResources []string) (*Turn, error) {
	return nil, nil
}
func (s *turnService) SetPlayerAction(gameId uuid.UUID, action Action) (*Turn, error) {
	return nil, nil
}
func (s *turnService) AddRandomEvent(gameId uuid.UUID, event Event) (*Turn, error) {
	return nil, nil
}
func (s *turnService) UpdateResources(gameId uuid.UUID, finalResources []string) (*Turn, error) {
	return nil, nil
}
