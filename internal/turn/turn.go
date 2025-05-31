// - The game is turn-based.
// - Each turn represents a year.
// - The player can only choose two actions per turn.

// Turns are divided into 3 actions:

// 1. Collect resources.
// 2. Player action.
//    a. Build.
//    b. Upgrade.
//    c. Allocate workers.
//    d. Collect taxes.
// 3. Events.

package turn

import "github.com/google/uuid"

type Action string
type Event string

const (
	Tax      Action = "collect"
	Build    Action = "build"
	Allocate Action = "allocate"
)

type Turn struct {
	Id               uuid.UUID
	Number           int
	PlayersAction    *Action
	RandomEvents     []Event
	InitialResources []string
	FinalResources   []string
}

func CreateFirstTurn() *Turn {
	return &Turn{
		Number:           1,
		PlayersAction:    nil,
		InitialResources: []string{},
		FinalResources:   []string{},
		RandomEvents:     []Event{},
	}
}

func (prev *Turn) CreateNextTurn() *Turn {
	return &Turn{
		Number:           prev.Number + 1,
		PlayersAction:    nil,
		InitialResources: prev.FinalResources,
		FinalResources:   []string{},
		RandomEvents:     []Event{},
	}
}

func (turn *Turn) SetPlayersAction(action *Action) *Turn {
	turn.PlayersAction = action
	return turn
}
