package goal

type Adapter struct {
}

func NewAdpater() *Adapter {
	return &Adapter{}
}
func (goa Adapter) SetGoal(title string, purpose1 string, purpose2 string) (int32, error) {
	return 1, nil
}

func (goa Adapter) GetGoal(goalId int32) (string, string, string, error) {
	return "hello", "Hello", "hello", nil
}
