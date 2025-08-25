package turn

import (
	"testing"

	"github.com/google/uuid"
)

func TestCreateTurn(t *testing.T) {
	t.Run("creates turn with correct initial state", func(t *testing.T) {
		gameID := uuid.New()
		turn := CreateTurn(gameID)

		if turn.GetID() == uuid.Nil {
			t.Error("expected turn ID to be generated")
		}
		if turn.GetGameID() != gameID {
			t.Errorf("expected gameID to be %s, got %s", gameID, turn.GetGameID())
		}
		if turn.GetNumber() != 1 {
			t.Errorf("expected turn number to be 1, got %d", turn.GetNumber())
		}
		if turn.GetStatus() != TurnStatusInProgress {
			t.Errorf("expected status to be %s, got %s", TurnStatusInProgress, turn.GetStatus())
		}
		if turn.GetCurrentPhase() != PhaseCollect {
			t.Errorf("expected phase to be %s, got %s", PhaseCollect, turn.GetCurrentPhase())
		}
		if turn.GetActionsUsed() != 0 {
			t.Errorf("expected actions used to be 0, got %d", turn.GetActionsUsed())
		}
		if turn.GetActionsRemaining() != MaxActionsPerTurn {
			t.Errorf("expected actions remaining to be %d, got %d", MaxActionsPerTurn, turn.GetActionsRemaining())
		}
		if turn.AreResourcesCollected() {
			t.Error("expected resources collected to be false")
		}
		if turn.AreEventsProcessed() {
			t.Error("expected events processed to be false")
		}
	})

	t.Run("each created turn has unique ID", func(t *testing.T) {
		gameID := uuid.New()
		turn1 := CreateTurn(gameID)
		turn2 := CreateTurn(gameID)

		if turn1.GetID() == turn2.GetID() {
			t.Error("expected different turn IDs")
		}
	})
}

func TestLoadTurn(t *testing.T) {
	t.Run("loads turn with all provided values", func(t *testing.T) {
		id := uuid.New()
		gameID := uuid.New()
		number := 5
		status := TurnStatusCompleted
		phase := PhaseEvents
		actionsUsed := 2
		resourcesCollected := true
		eventsProcessed := true

		turn := LoadTurn(id, gameID, number, status, phase, actionsUsed, resourcesCollected, eventsProcessed)

		if turn.GetID() != id {
			t.Errorf("expected ID to be %s, got %s", id, turn.GetID())
		}
		if turn.GetGameID() != gameID {
			t.Errorf("expected gameID to be %s, got %s", gameID, turn.GetGameID())
		}
		if turn.GetNumber() != number {
			t.Errorf("expected number to be %d, got %d", number, turn.GetNumber())
		}
		if turn.GetStatus() != status {
			t.Errorf("expected status to be %s, got %s", status, turn.GetStatus())
		}
		if turn.GetCurrentPhase() != phase {
			t.Errorf("expected phase to be %s, got %s", phase, turn.GetCurrentPhase())
		}
		if turn.GetActionsUsed() != actionsUsed {
			t.Errorf("expected actions used to be %d, got %d", actionsUsed, turn.GetActionsUsed())
		}
		if turn.AreResourcesCollected() != resourcesCollected {
			t.Errorf("expected resources collected to be %v, got %v", resourcesCollected, turn.AreResourcesCollected())
		}
		if turn.AreEventsProcessed() != eventsProcessed {
			t.Errorf("expected events processed to be %v, got %v", eventsProcessed, turn.AreEventsProcessed())
		}
	})

	t.Run("loads turn with minimal values", func(t *testing.T) {
		id := uuid.New()
		gameID := uuid.New()

		turn := LoadTurn(id, gameID, 1, TurnStatusInProgress, PhaseCollect, 0, false, false)

		if turn.GetID() != id {
			t.Errorf("expected ID to be %s, got %s", id, turn.GetID())
		}
		if turn.GetGameID() != gameID {
			t.Errorf("expected gameID to be %s, got %s", gameID, turn.GetGameID())
		}
	})
}

func TestCreateNextTurn(t *testing.T) {
	t.Run("creates next turn with incremented number", func(t *testing.T) {
		gameID := uuid.New()
		currentTurn := LoadTurn(uuid.New(), gameID, 5, TurnStatusCompleted, PhaseEvents, 2, true, true)
		
		nextTurn := currentTurn.CreateNextTurn()

		if nextTurn.GetID() == currentTurn.GetID() {
			t.Error("expected different ID for next turn")
		}
		if nextTurn.GetGameID() != gameID {
			t.Errorf("expected same gameID %s, got %s", gameID, nextTurn.GetGameID())
		}
		if nextTurn.GetNumber() != 6 {
			t.Errorf("expected turn number to be 6, got %d", nextTurn.GetNumber())
		}
		if nextTurn.GetStatus() != TurnStatusInProgress {
			t.Errorf("expected status to be %s, got %s", TurnStatusInProgress, nextTurn.GetStatus())
		}
		if nextTurn.GetCurrentPhase() != PhaseCollect {
			t.Errorf("expected phase to be %s, got %s", PhaseCollect, nextTurn.GetCurrentPhase())
		}
		if nextTurn.GetActionsUsed() != 0 {
			t.Errorf("expected actions used to be 0, got %d", nextTurn.GetActionsUsed())
		}
		if nextTurn.AreResourcesCollected() {
			t.Error("expected resources collected to be false")
		}
		if nextTurn.AreEventsProcessed() {
			t.Error("expected events processed to be false")
		}
	})

	t.Run("preserves gameID across turns", func(t *testing.T) {
		gameID := uuid.New()
		turn1 := CreateTurn(gameID)
		turn2 := turn1.CreateNextTurn()
		turn3 := turn2.CreateNextTurn()

		if turn1.GetGameID() != gameID || turn2.GetGameID() != gameID || turn3.GetGameID() != gameID {
			t.Error("expected all turns to have same gameID")
		}
		if turn1.GetNumber() != 1 || turn2.GetNumber() != 2 || turn3.GetNumber() != 3 {
			t.Error("expected sequential turn numbers: 1, 2, 3")
		}
	})
}

func TestGetActionsRemaining(t *testing.T) {
	testCases := []struct {
		name            string
		actionsUsed     int
		expectedRemaining int
	}{
		{"no actions used", 0, 2},
		{"one action used", 1, 1},
		{"max actions used", 2, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gameID := uuid.New()
			turn := LoadTurn(uuid.New(), gameID, 1, TurnStatusInProgress, PhaseAction, tc.actionsUsed, false, false)

			remaining := turn.GetActionsRemaining()
			if remaining != tc.expectedRemaining {
				t.Errorf("expected %d actions remaining, got %d", tc.expectedRemaining, remaining)
			}
		})
	}
}

func TestTurnStatusConstants(t *testing.T) {
	t.Run("constants have expected values", func(t *testing.T) {
		if TurnStatusInProgress != "in_progress" {
			t.Errorf("expected TurnStatusInProgress to be 'in_progress', got %s", TurnStatusInProgress)
		}
		if TurnStatusCompleted != "completed" {
			t.Errorf("expected TurnStatusCompleted to be 'completed', got %s", TurnStatusCompleted)
		}
	})
}

func TestTurnPhaseConstants(t *testing.T) {
	t.Run("constants have expected values", func(t *testing.T) {
		if PhaseCollect != "collect" {
			t.Errorf("expected PhaseCollect to be 'collect', got %s", PhaseCollect)
		}
		if PhaseAction != "action" {
			t.Errorf("expected PhaseAction to be 'action', got %s", PhaseAction)
		}
		if PhaseEvents != "events" {
			t.Errorf("expected PhaseEvents to be 'events', got %s", PhaseEvents)
		}
	})
}

func TestTurnStateTransitions(t *testing.T) {
	t.Run("turn progresses through expected phases", func(t *testing.T) {
		gameID := uuid.New()
		
		// Test collect phase
		turn := CreateTurn(gameID)
		if turn.GetCurrentPhase() != PhaseCollect {
			t.Errorf("expected initial phase to be %s, got %s", PhaseCollect, turn.GetCurrentPhase())
		}
		
		// Test action phase
		actionTurn := LoadTurn(uuid.New(), gameID, 1, TurnStatusInProgress, PhaseAction, 1, true, false)
		if actionTurn.GetCurrentPhase() != PhaseAction {
			t.Errorf("expected phase to be %s, got %s", PhaseAction, actionTurn.GetCurrentPhase())
		}
		if actionTurn.GetActionsRemaining() != 1 {
			t.Errorf("expected 1 action remaining, got %d", actionTurn.GetActionsRemaining())
		}
		
		// Test events phase
		eventsTurn := LoadTurn(uuid.New(), gameID, 1, TurnStatusInProgress, PhaseEvents, 2, true, false)
		if eventsTurn.GetCurrentPhase() != PhaseEvents {
			t.Errorf("expected phase to be %s, got %s", PhaseEvents, eventsTurn.GetCurrentPhase())
		}
		if eventsTurn.GetActionsRemaining() != 0 {
			t.Errorf("expected 0 actions remaining, got %d", eventsTurn.GetActionsRemaining())
		}
		
		// Test completed turn
		completedTurn := LoadTurn(uuid.New(), gameID, 1, TurnStatusCompleted, PhaseEvents, 2, true, true)
		if completedTurn.GetStatus() != TurnStatusCompleted {
			t.Errorf("expected status to be %s, got %s", TurnStatusCompleted, completedTurn.GetStatus())
		}
		if !completedTurn.AreEventsProcessed() {
			t.Error("expected events to be processed")
		}
	})
}