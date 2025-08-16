package resource

import (
	"testing"
)

func TestNewResource(t *testing.T) {
	t.Run("new resource with value", func(t *testing.T) {
		value := 10
		resource := NewResource(value)
		if resource.GetValue() != value {
			t.Errorf("Expected value to be equal %d", value)
		}
	})
	t.Run("new resource empty", func(t *testing.T) {
		value := 0
		resource := NewResource()
		if resource.GetValue() != value {
			t.Errorf("Expected value to be equal %d", value)
		}
	})
	t.Run("new resource with multiple values should return the first", func(t *testing.T) {
		value := 10
		resource := NewResource(value, 11, 12, 13)
		if resource.GetValue() != value {
			t.Errorf("Expected value to be equal %d", value)
		}
	})
}

func TestAdjustResourceValue(t *testing.T) {
	testCases := []struct {
		testName string
		initialValue int
		adjustment int
		expectedValue int
		expectError bool
	} {
		{"positive adjustment", 10, 5, 15, false},
        {"negative valid adjustment", 10, -5, 5, false},
        {"invalid negative result", 5, -10, 5, true},
        {"zero adjustment", 10, 0, 10, false},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			resource := NewResource(tc.initialValue)
			_, err := resource.AdjustValue(tc.adjustment)

			if tc.expectError && err == nil {
                t.Fatal("expected error but got none")
            }
            if !tc.expectError && err != nil {
                t.Fatalf("unexpected error: %v", err)
            }
            if resource.GetValue() != tc.expectedValue {
                t.Fatalf("expected %d, got %d", tc.expectedValue, resource.GetValue())
            }
		})
	}
}

func TestIsOperationValid(t *testing.T) {
	testCases := []struct {
		testName     string
		initialValue int
		adjustment   int
		expectError  bool
	}{
		{"valid positive adjustment", 10, 5, false},
		{"valid negative adjustment", 10, -5, false},
		{"valid zero adjustment", 10, 0, false},
		{"invalid negative result", 5, -10, true},
		{"boundary case - exactly zero result", 5, -5, false},
		{"boundary case - exactly negative by one", 5, -6, true},
		{"zero initial value with positive adjustment", 0, 1, false},
		{"zero initial value with negative adjustment", 0, -1, true},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			resource := NewResource(tc.initialValue)
			err := resource.IsOperationValid(tc.adjustment)

			if tc.expectError && err == nil {
				t.Fatal("expected error but got none")
			}
			if !tc.expectError && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			
			if resource.GetValue() != tc.initialValue {
				t.Fatalf("validation changed resource value from %d to %d", tc.initialValue, resource.GetValue())
			}
		})
	}
}