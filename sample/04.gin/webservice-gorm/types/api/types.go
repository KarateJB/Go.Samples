package types

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type TrackDateTimes struct {
	CreateOn time.Time    `json:"createOn"`
	UpdateOn sql.NullTime `json:"updateOn"`
	DeleteOn sql.NullTime `json:"deleteOn"`
}

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Todos []Todo `json:"todos"`
}

type Todo struct {
	Id             uuid.UUID `json:"id"`
	Title          string    `json:"title"`
	IsDone         bool      `json:"isDone"`
	TrackDateTimes `json:"trackDateTimes"`
	TodoExt        TodoExt `json:"todoExt"`
	UserId         string  `json:"userId"`
	Tags           []Tag   `json:"tags"`
}

type TodoExt struct {
	Id          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	Priority    Priority  `json:""`
}

type Tag struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type Priority struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// TODO: Remove this struct
type TodoPageData struct {
	PageTitle string `json:"pageTitle"`
	Todos     []Todo `json:"todos"`
}
