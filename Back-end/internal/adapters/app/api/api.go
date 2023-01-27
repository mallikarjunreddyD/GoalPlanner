package api

import "github.com/mallikarjun/GoalPlanner/Back-end/internal/ports"

type Adapter struct {
	goal ports.GoalPort
	db   ports.DbPort
}

func NewAdpater(db ports.DbPort, goal ports.GoalPort) *Adapter {
	return &Adapter{db: db, goal: goal}
}
func (apia Adapter) SetGoal(title string, purpose1 string, purpose2 string) error {
	err := apia.goal.SetGoal(title, purpose1, purpose2)
	if err != nil {
		return err
	}
	err = apia.db.AddNewGoal(title, purpose1, purpose2)
	if err != nil {
		return err
	}
	return nil
}
func (apia Adapter) GetGoal(goalId int32) (string, string, string, error) {
	title, purpose1, purpose2, err := apia.goal.GetGoal(goalId)
	if err != nil {
		return "", "", "", err
	}
	return title, purpose1, purpose2, nil
}
