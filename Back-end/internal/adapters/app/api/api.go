package api

import (
	"github.com/mallikarjun/GoalPlanner/Back-end/internal/ports"
)

type Adapter struct {
	goal ports.GoalPort
}

func NewAdpater(goal ports.GoalPort) *Adapter {
	return &Adapter{goal: goal}
}
func (apia Adapter) SetGoal(title string, purpose1 string, purpose2 string) (int32, error) {
	id, err := apia.goal.SetGoal(title, purpose1, purpose2)
	return id, err
}
