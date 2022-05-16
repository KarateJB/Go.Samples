package types

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Id     uuid.UUID `gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	Title  string
	IsDone bool
}

// TableName: Specified table name for struct Todo
func (Todo) TableName() string {
	return "Todos"
}
