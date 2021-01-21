package types

// Todo is a todo task
type Todo struct {
	Title  string
	IsDone bool
}

// TodoPageData contains todo title and tasks
type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}
