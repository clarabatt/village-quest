package game

import (
	"testing"
)

var A_GAME_NUMBER = 77
var A_PLAYER_NAME = "John Doe"

func TestNewGame(t *testing.T) {
	game := NewGame(A_GAME_NUMBER, A_PLAYER_NAME)

	if game.number != A_GAME_NUMBER {
		t.Errorf("Expected game.Number to be %d, got %d", A_GAME_NUMBER, game.number)
	}

	if game.playersName != A_PLAYER_NAME {
		t.Errorf("Expected game.playersName to be %s, got %s", A_PLAYER_NAME, game.playersName)
	}

	if game.maxDaysPlayed != 0 {
		t.Errorf("Expected game.maxDaysPlayed to be 0, got %d", game.maxDaysPlayed)
	}
}
