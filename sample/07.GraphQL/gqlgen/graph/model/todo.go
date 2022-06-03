package model

import "github.com/google/uuid"

type Todo struct {
	Id     uuid.UUID `json:"id"`
	Title  string    `json:"title"`
	IsDone bool      `json:"isDone"`
	UserId string    `json:"userId"`
	User   *User     `json:"user"`
}

type NewTodo struct {
	Title  string `json:"title"`
	IsDone bool   `json:"isDone"`
	UserId string `json:"userId"`
}

type EditTodo struct {
	Id     uuid.UUID `json:"id"`
	Title  string    `json:"title"`
	IsDone bool      `json:"isDone"`
}
