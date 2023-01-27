package ports

type DbPort interface {
	CloseDbConnection()
	AddNewGoal(title string, purpose1 string, purpose2 string) error
}
