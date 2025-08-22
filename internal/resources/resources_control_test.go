package resources

import (
	"testing"
)

func TestNewResourceControl(t *testing.T) {
	t.Run("creates resource control with specified values", func(t *testing.T) {
		rc := NewResourceControl(10, 5, 3, 8, 2)
		
		if rc.GetStone() != 10 {
			t.Errorf("expected stone to be 10, got %d", rc.GetStone())
		}
		if rc.GetGold() != 5 {
			t.Errorf("expected gold to be 5, got %d", rc.GetGold())
		}
		if rc.GetWood() != 3 {
			t.Errorf("expected wood to be 3, got %d", rc.GetWood())
		}
		if rc.GetFood() != 8 {
			t.Errorf("expected food to be 8, got %d", rc.GetFood())
		}
		if rc.GetWorker() != 2 {
			t.Errorf("expected worker to be 2, got %d", rc.GetWorker())
		}
	})

	t.Run("creates resource control with zero values", func(t *testing.T) {
		rc := NewResourceControl(0, 0, 0, 0, 0)
		
		if rc.GetStone() != 0 {
			t.Errorf("expected stone to be 0, got %d", rc.GetStone())
		}
		if rc.GetGold() != 0 {
			t.Errorf("expected gold to be 0, got %d", rc.GetGold())
		}
		if rc.GetWood() != 0 {
			t.Errorf("expected wood to be 0, got %d", rc.GetWood())
		}
		if rc.GetFood() != 0 {
			t.Errorf("expected food to be 0, got %d", rc.GetFood())
		}
		if rc.GetWorker() != 0 {
			t.Errorf("expected worker to be 0, got %d", rc.GetWorker())
		}
	})
}

func TestGetResourcesMap(t *testing.T) {
	t.Run("returns correct resource map", func(t *testing.T) {
		rc := NewResourceControl(10, 5, 3, 8, 2)
		resourceMap := rc.GetResourcesMap()
		
		expected := map[string]int{
			"Stone":  10,
			"Gold":   5,
			"Wood":   3,
			"Food":   8,
			"Worker": 2,
		}
		
		for key, expectedValue := range expected {
			if resourceMap[key] != expectedValue {
				t.Errorf("expected %s to be %d, got %d", key, expectedValue, resourceMap[key])
			}
		}
	})
}

func TestIndividualAdjustMethods(t *testing.T) {
	testCases := []struct {
		testName     string
		adjustMethod func(*ResourcesControl, int) (int, error)
		getMethod    func(*ResourcesControl) int
		resourceName string
	}{
		{"stone", (*ResourcesControl).AdjustStone, (*ResourcesControl).GetStone, "stone"},
		{"gold", (*ResourcesControl).AdjustGold, (*ResourcesControl).GetGold, "gold"},
		{"wood", (*ResourcesControl).AdjustWood, (*ResourcesControl).GetWood, "wood"},
		{"food", (*ResourcesControl).AdjustFood, (*ResourcesControl).GetFood, "food"},
		{"worker", (*ResourcesControl).AdjustWorker, (*ResourcesControl).GetWorker, "worker"},
	}

	for _, tc := range testCases {
		t.Run("adjust "+tc.resourceName+" positive", func(t *testing.T) {
			rc := NewResourceControl(10, 10, 10, 10, 10)
			newValue, err := tc.adjustMethod(rc, 5)
			
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if newValue != 15 {
				t.Errorf("expected returned value to be 15, got %d", newValue)
			}
			if tc.getMethod(rc) != 15 {
				t.Errorf("expected %s to be 15, got %d", tc.resourceName, tc.getMethod(rc))
			}
		})

		t.Run("adjust "+tc.resourceName+" negative valid", func(t *testing.T) {
			rc := NewResourceControl(10, 10, 10, 10, 10)
			newValue, err := tc.adjustMethod(rc, -3)
			
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if newValue != 7 {
				t.Errorf("expected returned value to be 7, got %d", newValue)
			}
			if tc.getMethod(rc) != 7 {
				t.Errorf("expected %s to be 7, got %d", tc.resourceName, tc.getMethod(rc))
			}
		})

		t.Run("adjust "+tc.resourceName+" negative invalid", func(t *testing.T) {
			rc := NewResourceControl(5, 5, 5, 5, 5)
			_, err := tc.adjustMethod(rc, -10)
			
			if err == nil {
				t.Fatal("expected error for invalid negative adjustment")
			}
			if tc.getMethod(rc) != 5 {
				t.Errorf("expected %s to remain 5 after failed adjustment, got %d", tc.resourceName, tc.getMethod(rc))
			}
		})
	}
}

func TestAdjustMultiple(t *testing.T) {
	t.Run("successful multiple adjustments", func(t *testing.T) {
		rc := NewResourceControl(10, 10, 10, 10, 10)
		adjustments := map[string]int{
			"stone": 5,
			"gold":  -3,
			"wood":  0,
			"food":  2,
			"worker": -1,
		}
		
		err := rc.AdjustMultiple(adjustments)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		
		expected := map[string]int{
			"stone":  15,
			"gold":   7,
			"wood":   10,
			"food":   12,
			"worker": 9,
		}
		
		if rc.GetStone() != expected["stone"] {
			t.Errorf("expected stone to be %d, got %d", expected["stone"], rc.GetStone())
		}
		if rc.GetGold() != expected["gold"] {
			t.Errorf("expected gold to be %d, got %d", expected["gold"], rc.GetGold())
		}
		if rc.GetWood() != expected["wood"] {
			t.Errorf("expected wood to be %d, got %d", expected["wood"], rc.GetWood())
		}
		if rc.GetFood() != expected["food"] {
			t.Errorf("expected food to be %d, got %d", expected["food"], rc.GetFood())
		}
		if rc.GetWorker() != expected["worker"] {
			t.Errorf("expected worker to be %d, got %d", expected["worker"], rc.GetWorker())
		}
	})

	t.Run("case insensitive resource names", func(t *testing.T) {
		rc := NewResourceControl(10, 10, 10, 10, 10)
		adjustments := map[string]int{
			"STONE": 5,
			"Gold":  3,
			"wOoD":  2,
		}
		
		err := rc.AdjustMultiple(adjustments)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		
		if rc.GetStone() != 15 {
			t.Errorf("expected stone to be 15, got %d", rc.GetStone())
		}
		if rc.GetGold() != 13 {
			t.Errorf("expected gold to be 13, got %d", rc.GetGold())
		}
		if rc.GetWood() != 12 {
			t.Errorf("expected wood to be 12, got %d", rc.GetWood())
		}
	})

	t.Run("unknown resource name", func(t *testing.T) {
		rc := NewResourceControl(10, 10, 10, 10, 10)
		adjustments := map[string]int{
			"unknown": 5,
		}
		
		err := rc.AdjustMultiple(adjustments)
		if err == nil {
			t.Fatal("expected error for unknown resource")
		}
		
		expectedError := "unknown resource: unknown"
		if err.Error() != expectedError {
			t.Errorf("expected error '%s', got '%s'", expectedError, err.Error())
		}
	})

	t.Run("invalid adjustment - atomic behavior", func(t *testing.T) {
		rc := NewResourceControl(5, 10, 15, 20, 25)
		originalValues := map[string]int{
			"stone":  rc.GetStone(),
			"gold":   rc.GetGold(),
			"wood":   rc.GetWood(),
			"food":   rc.GetFood(),
			"worker": rc.GetWorker(),
		}
		
		adjustments := map[string]int{
			"stone": 5,   // valid: 5 + 5 = 10
			"gold":  -15, // invalid: 10 - 15 = -5
			"wood":  3,   // valid: 15 + 3 = 18
		}
		
		err := rc.AdjustMultiple(adjustments)
		if err == nil {
			t.Fatal("expected error for invalid adjustment")
		}
		
		if rc.GetStone() != originalValues["stone"] {
			t.Errorf("expected stone to remain %d, got %d", originalValues["stone"], rc.GetStone())
		}
		if rc.GetGold() != originalValues["gold"] {
			t.Errorf("expected gold to remain %d, got %d", originalValues["gold"], rc.GetGold())
		}
		if rc.GetWood() != originalValues["wood"] {
			t.Errorf("expected wood to remain %d, got %d", originalValues["wood"], rc.GetWood())
		}
		if rc.GetFood() != originalValues["food"] {
			t.Errorf("expected food to remain %d, got %d", originalValues["food"], rc.GetFood())
		}
		if rc.GetWorker() != originalValues["worker"] {
			t.Errorf("expected worker to remain %d, got %d", originalValues["worker"], rc.GetWorker())
		}
	})

	t.Run("empty adjustments map", func(t *testing.T) {
		rc := NewResourceControl(10, 10, 10, 10, 10)
		originalValues := map[string]int{
			"stone":  rc.GetStone(),
			"gold":   rc.GetGold(),
			"wood":   rc.GetWood(),
			"food":   rc.GetFood(),
			"worker": rc.GetWorker(),
		}
		
		err := rc.AdjustMultiple(map[string]int{})
		if err != nil {
			t.Fatalf("unexpected error for empty map: %v", err)
		}
		
		if rc.GetStone() != originalValues["stone"] {
			t.Errorf("expected stone to remain %d, got %d", originalValues["stone"], rc.GetStone())
		}
		if rc.GetGold() != originalValues["gold"] {
			t.Errorf("expected gold to remain %d, got %d", originalValues["gold"], rc.GetGold())
		}
	})
}