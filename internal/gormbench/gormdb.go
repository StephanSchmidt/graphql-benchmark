package gormbench

type GormTaskList struct {
	Tasks []Task
}

type User struct {
	ID   uint
	Name string
}

type Status struct {
	ID     uint
	Status string
}

type Task struct {
	ID       uint
	Title    string
	UserID   uint
	User     User
	StatusID uint
	Status   Status
}
