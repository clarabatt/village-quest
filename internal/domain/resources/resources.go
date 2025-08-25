package resources

import (
	"fmt"
	"strings"
)

type Resources struct {
	stone  *Resource
	gold   *Resource
	wood   *Resource
	food   *Resource
	worker *Resource
}

func NewResourceControl(stone int, gold int, wood int, food int, worker int) *Resources {
	r := Resources{
		stone:  NewResource(stone),
		gold:   NewResource(gold),
		wood:   NewResource(wood),
		food:   NewResource(food),
		worker: NewResource(worker),
	}
	return &r
}

func (rc *Resources) GetResourcesMap() map[string]int {
	return map[string]int{
		"Stone":  rc.GetStone(),
		"Gold":   rc.GetGold(),
		"Wood":   rc.GetWood(),
		"Food":   rc.GetFood(),
		"Worker": rc.GetWorker(),
	}
}

func (rc *Resources) AdjustMultiple(adjustments map[string]int) error {
	for resourceName, adjustment := range adjustments {
		var currentResource *Resource

		switch strings.ToLower(resourceName) {
		case "stone":
			currentResource = rc.stone
		case "gold":
			currentResource = rc.gold
		case "wood":
			currentResource = rc.wood
		case "food":
			currentResource = rc.food
		case "worker":
			currentResource = rc.worker
		default:
			return fmt.Errorf("unknown resource: %s", resourceName)
		}

		if err := currentResource.IsOperationValid(adjustment); err != nil {
			return fmt.Errorf("invalid adjustment for %s: %w", resourceName, err)
		}
	}

	for resourceName, adjustment := range adjustments {
		switch strings.ToLower(resourceName) {
		case "stone":
			rc.stone.AdjustValue(adjustment)
		case "gold":
			rc.gold.AdjustValue(adjustment)
		case "wood":
			rc.wood.AdjustValue(adjustment)
		case "food":
			rc.food.AdjustValue(adjustment)
		case "worker":
			rc.worker.AdjustValue(adjustment)
		}
	}
	return nil
}

func (rc *Resources) GetStone() int {
	return rc.stone.GetValue()
}

func (rc *Resources) GetGold() int {
	return rc.gold.GetValue()
}

func (rc *Resources) GetWood() int {
	return rc.wood.GetValue()
}

func (rc *Resources) GetFood() int {
	return rc.food.GetValue()
}

func (rc *Resources) GetWorker() int {
	return rc.worker.GetValue()
}

func (rc *Resources) AdjustStone(value int) (int, error) {
	return rc.stone.AdjustValue(value)
}

func (rc *Resources) AdjustGold(value int) (int, error) {
	return rc.gold.AdjustValue(value)
}

func (rc *Resources) AdjustWood(value int) (int, error) {
	return rc.wood.AdjustValue(value)
}

func (rc *Resources) AdjustFood(value int) (int, error) {
	return rc.food.AdjustValue(value)
}

func (rc *Resources) AdjustWorker(value int) (int, error) {
	return rc.worker.AdjustValue(value)
}
