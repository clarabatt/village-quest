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
	t.Run("should increase if value is positive and valid", func(t *testing.T) {
		startValue := 10
		resource := NewResource(startValue)
		increasedValue := 20
		newValue, err := resource.AdjustValue(increasedValue)
		expectedValue := startValue + increasedValue
		if newValue != expectedValue {
			t.Errorf("Expected value to be equal %d + %d = %d, got %d", startValue, increasedValue, expectedValue, newValue)
		}
		if resource.GetValue() != expectedValue {
			t.Errorf("Expected value to be equal %d + %d = %d, got %d", startValue, increasedValue, expectedValue, resource.GetValue())
		}
		if err != nil {
			t.Error("expected no error, got:", err)
		}
	})

	t.Run("should return error if value would result in a negative number", func(t *testing.T) {
		startValue := 10
		resource := NewResource(startValue)
		increasedValue := -50
		newValue, err := resource.AdjustValue(increasedValue)
		expectedValue := startValue
		if newValue != 0 {
			t.Errorf("Expected value to be 0, got %d", newValue)
		}
		if resource.GetValue() != expectedValue {
			t.Errorf("Expected resource value to not be changed")
		}
		if err == nil {
			t.Errorf("Expected error for invalid value, got nil")
		}
	})

}