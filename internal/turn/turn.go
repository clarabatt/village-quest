package turn

import "github.com/google/uuid"

type TurnStatus string
type TurnPhase string

const (
    TurnStatusInProgress TurnStatus = "in_progress"
    TurnStatusCompleted  TurnStatus = "completed"
)
const (
    PhaseCollect TurnPhase = "collect"
    PhaseAction  TurnPhase = "action" 
    PhaseEvents  TurnPhase = "events"
)

const MaxActionsPerTurn = 2

type Turn struct {
	id               uuid.UUID
	gameID uuid.UUID
	number           int
	status TurnStatus
	currentPhase     TurnPhase
	actionsUsed      int
	resourcesCollected bool
	eventsProcessed   bool
}

func CreateTurn(gameID uuid.UUID) *Turn {
	return &Turn{
		id:                 uuid.New(),
		gameID:             gameID,
		number:             1,
		status:             TurnStatusInProgress,
		currentPhase:       PhaseCollect,
		actionsUsed:        0,
		resourcesCollected: false,
		eventsProcessed:    false,
	}
}

func LoadTurn(id uuid.UUID, gameID uuid.UUID, number int, status TurnStatus, 
	currentPhase TurnPhase, actionsUsed int, resourcesCollected bool, eventsProcessed bool) *Turn {
	return &Turn{
		id:                 id,
		gameID:             gameID,
		number:             number,
		status:             status,
		currentPhase:       currentPhase,
		actionsUsed:        actionsUsed,
		resourcesCollected: resourcesCollected,
		eventsProcessed:    eventsProcessed,
	}
}

func (t *Turn) CreateNextTurn() *Turn {	
	return &Turn{
		id:                 uuid.New(),
		gameID:             t.gameID,
		number:             t.number + 1,
		status:             TurnStatusInProgress,
		currentPhase:       PhaseCollect,
		actionsUsed:        0,
		resourcesCollected: false,
		eventsProcessed:    false,
	}
}

func (t *Turn) GetID() uuid.UUID {
	return t.id
}

func (t *Turn) GetGameID() uuid.UUID {
	return t.gameID
}

func (t *Turn) GetNumber() int {
	return t.number
}

func (t *Turn) GetStatus() TurnStatus {
	return t.status
}

func (t *Turn) GetCurrentPhase() TurnPhase {
	return t.currentPhase
}

func (t *Turn) GetActionsUsed() int {
	return t.actionsUsed
}

func (t *Turn) GetActionsRemaining() int {
	return MaxActionsPerTurn - t.actionsUsed
}

func (t *Turn) AreResourcesCollected() bool {
	return t.resourcesCollected
}

func (t *Turn) AreEventsProcessed() bool {
	return t.eventsProcessed
}