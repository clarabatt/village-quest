package game

import (
	"testing"

	"github.com/google/uuid"
)

var A_GAME_ID = uuid.New()
var A_GAME_DAYS_PLAYED = 40
var A_GAME_NUMBER = 77
var A_PLAYER_NAME = "John Doe"

func TestCreateGame(t *testing.T) {
	t.Run("New Game Creation", func(t *testing.T) {
		game := NewGame(A_GAME_NUMBER, A_PLAYER_NAME)

		if game.number != A_GAME_NUMBER {
			t.Errorf("Expected game.number to be %d, got %d", A_GAME_NUMBER, game.number)
		}
		if game.playersName != A_PLAYER_NAME {
			t.Errorf("Expected game.playersName to be %s, got %s", A_PLAYER_NAME, game.playersName)
		}
		if game.daysPlayed != 0 {
			t.Errorf("Expected game.daysPlayed to be 0, got %d", game.daysPlayed)
		}
	})
}
func TestLoadGame(t *testing.T) {
	t.Run("Load Existing Game", func(t *testing.T) {
		game := LoadGame(A_GAME_ID, A_GAME_NUMBER, A_GAME_DAYS_PLAYED, A_PLAYER_NAME)

		if game.id != A_GAME_ID {
			t.Errorf("Expected game.id to be %d, got %d", A_GAME_ID, game.id)
		}
		if game.number != A_GAME_NUMBER {
			t.Errorf("Expected game.number to be %d, got %d", A_GAME_NUMBER, game.number)
		}
		if game.playersName != A_PLAYER_NAME {
			t.Errorf("Expected game.playersName to be %s, got %s", A_PLAYER_NAME, game.playersName)
		}
		if game.daysPlayed != A_GAME_DAYS_PLAYED {
			t.Errorf("Expected game.daysPlayed to be %d, got %d", A_GAME_DAYS_PLAYED, game.daysPlayed)
		}
	})
}
func TestGameGetters(t *testing.T) {
	t.Run("Getters", func(t *testing.T) {
		game := LoadGame(A_GAME_ID, A_GAME_NUMBER, A_GAME_DAYS_PLAYED, A_PLAYER_NAME)

		if game.Id() != A_GAME_ID {
			t.Errorf("Expected game.id to be %v, got %v", A_GAME_ID, game.Id())
		}
		if game.Number() != A_GAME_NUMBER {
			t.Errorf("Expected game.number to be %d, got %d", A_GAME_NUMBER, game.Number())
		}
		if game.PlayersName() != A_PLAYER_NAME {
			t.Errorf("Expected game.playersName to be %s, got %s", A_PLAYER_NAME, game.PlayersName())
		}
		if game.DaysPlayed() != A_GAME_DAYS_PLAYED {
			t.Errorf("Expected game.daysPlayed to be %d, got %d", A_GAME_DAYS_PLAYED, game.DaysPlayed())
		}
	})
}
