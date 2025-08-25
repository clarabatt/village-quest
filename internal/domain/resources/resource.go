package resources

import "fmt" 

type Resource struct {
	value int
}

func NewResource(values ...int) *Resource {
	if len(values) == 0 {
		return &Resource { value: 0 }
	}
	return &Resource { value: values[0] }
}

func (r *Resource) GetValue() int {
	return r.value 
}

func (r *Resource) IsOperationValid(value int) error {
	if r.value + value < 0 {
		return fmt.Errorf("operation would result in negative value: %d + %d = %d", 
    r.value, value, r.value+value)
	}
	return nil
}

func (r *Resource) AdjustValue(value int) (int, error){
	err := r.IsOperationValid(value) 
	if err != nil {
		return 0, err
	}
	r.value += value
	return r.value, nil
}

