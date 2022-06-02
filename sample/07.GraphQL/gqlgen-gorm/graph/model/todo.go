package model

import "github.com/google/uuid"

type Todo struct {
	Id      uuid.UUID `json:"id"`
	Title   string    `json:"title"`
	IsDone  bool      `json:"isDone"`
	TodoExt *TodoExt  `json:"todoExt"`
	UserId  string    `json:"userId"`
	User    *User     `json:"user"`
	Tags    []*Tag    `json:"tags"`
}

type NewTodo struct {
	Title   string      `json:"title"`
	IsDone  bool        `json:"isDone"`
	TodoExt *NewTodoExt `json:"todoExt"`
	UserId  string      `json:"userId"`
	TagIds  []uuid.UUID `json:"tags"`
}

type EditTodo struct {
	Id      uuid.UUID    `json:"id"`
	Title   string       `json:"title"`
	IsDone  bool         `json:"isDone"`
	TodoExt *EditTodoExt `json:"todoExt"`
	TagIds  []uuid.UUID  `json:"tags"`
}

type EditTodoExt struct {
	Id          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	PriorityId  int       `json:"priorityId"`
}
type NewTodoExt struct {
	Description string `json:"description"`
	PriorityId  int    `json:"priorityId"`
}

type Priority struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Tag struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type TodoExt struct {
	Id          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	PriorityId  int       `json:"priorityId"`
	Priority    *Priority `json:"priority"`
}
