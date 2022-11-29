package pgxbench

type PgxTaskList struct {
	Tasks []PgxTask
}

type PgxTask struct {
	Id     int64
	Title  string
	Status string
	User   string
}
