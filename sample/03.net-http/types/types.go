package types

// Todo is a todo task
type Todo struct {
	Title  string `json:"title"`
	IsDone bool   `json:"isDone"`
}

// TodoPageData contains todo title and tasks
type TodoPageData struct {
	PageTitle string `json:"pageTitle"`
	Todos     []Todo `json:"todos"`
}
