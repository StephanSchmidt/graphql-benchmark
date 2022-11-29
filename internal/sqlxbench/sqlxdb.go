package sqlxbench

type SqlxTaskList struct {
	Tasks []SqlxTask
}

type SqlxTask struct {
	Id     int64
	Title  string
	Status string
	Name   string
}
