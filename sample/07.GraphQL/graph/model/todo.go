package model

import "github.com/google/uuid"

type Todo struct {
	Id     uuid.UUID `json:"id"`
	Title  string    `json:"title"`
	IsDone bool      `json:"isDone"`
	User   *User     `json:"user"`
}

type NewTodo struct {
	Title  string `json:"title"`
	IsDone bool   `json:"isDone"`
	UserID string `json:"userId"`
}
