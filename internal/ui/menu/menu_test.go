package menu

import (
	"fmt"
	"testing"
)

func mockAction() {
	fmt.Print("Action!")
}

var A_TITLE = "New Menu"
var AN_EXIT_OPTION = 999

var AN_ITEM_NAME = "Item"
var AN_ITEM_ACTION = mockAction
var AN_ITEM_ORDER = 1

func TestCreateMenu(t *testing.T) {
	t.Run("without exit option", func(t *testing.T) {
		menu := NewMenu(A_TITLE, nil)

		if menu.Title != A_TITLE {
			t.Errorf("Expected Title to be %s, got %s", A_TITLE, menu.Title)
		}
		if menu.ExitOption != 0 {
			t.Errorf("Expected ExitOption to be %d, got %d", 0, menu.ExitOption)
		}
	})
	t.Run("with exit option provided", func(t *testing.T) {
		menu := NewMenu(A_TITLE, &AN_EXIT_OPTION)

		if menu.Title != A_TITLE {
			t.Errorf("Expected Title to be %s, got %s", A_TITLE, menu.Title)
		}
		if menu.ExitOption != AN_EXIT_OPTION {
			t.Errorf("Expected ExitOption to be %d, got %d", AN_EXIT_OPTION, menu.ExitOption)
		}
	})
}

func TestAddItem(t *testing.T) {
	t.Run("adds a valid item", func(t *testing.T) {
		menu := NewMenu(A_TITLE, nil)
		err := menu.AddItem(AN_ITEM_NAME, AN_ITEM_ACTION, AN_ITEM_ORDER)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if len(menu.Items) != 1 {
			t.Errorf("Expected 1 item in menu, got %d", len(menu.Items))
		}
		if menu.Items[0].Name != AN_ITEM_NAME {
			t.Errorf("Expected item name to be %s, got %s", AN_ITEM_NAME, menu.Items[0].Name)
		}
		if menu.Items[0].Order != AN_ITEM_ORDER {
			t.Errorf("Expected item order to be %d, got %d", AN_ITEM_ORDER, menu.Items[0].Order)
		}
	})

	t.Run("fails on duplicate order", func(t *testing.T) {
		menu := NewMenu(A_TITLE, nil)
		_ = menu.AddItem("Item 1", AN_ITEM_ACTION, 1)
		err := menu.AddItem("Item 2", AN_ITEM_ACTION, 1)

		if err == nil {
			t.Errorf("Expected error for duplicate order, got nil")
		}
	})

	t.Run("fails when order matches exit option", func(t *testing.T) {
		exit := 5
		menu := NewMenu(A_TITLE, &exit)
		err := menu.AddItem("ExitConflict", AN_ITEM_ACTION, 5)

		if err == nil {
			t.Errorf("Expected error for item with same order as exit option, got nil")
		}
	})
}
