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
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Todo struct {
	Id     uuid.UUID `json:"id"`
	Title  string    `json:"title"`
	IsDone bool      `json:"isDone"`
	// TrackDateTimes `json:"trackDateTimes"`
	TodoExt TodoExt `json:"todoExt"`
	UserId  string  `json:"userId"`
	User    User    `json:"user"`
	Tags    []Tag   `json:"tags"`
}

type TodoExt struct {
	Id          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	PriorityId  int       `json:"priorityId"`
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
