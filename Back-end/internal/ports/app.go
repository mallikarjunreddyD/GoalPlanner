package ports

type APIPort interface {
	SetGoal(title string, purpose1 string, purpose2 string) (int32, error)
	GetGoal(goalId int32) (string, string, string, error)
}
