package types

type Todo struct {
	Title  string `json:"title"`
	IsDone bool   `json:"isDone"`
}

type TodoPageData struct {
	PageTitle string `json:"pageTitle"`
	Todos     []Todo `json:"todos"`
}
