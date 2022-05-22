package types

import "github.com/google/uuid"

type Todo struct {
	Id     uuid.UUID `json:"id"`
	Title  string    `json:"title"`
	IsDone bool      `json:"isDone"`
}

type TodoPageData struct {
	PageTitle string `json:"pageTitle"`
	Todos     []Todo `json:"todos"`
}
