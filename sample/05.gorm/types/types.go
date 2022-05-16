package types

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type TrackDateTimes struct {
	CreateOn time.Time    `gorm:"column:CreateOn;not null;default:now();comment:Created datetime(UTC)"`
	UpdateOn sql.NullTime `gorm:"column:UpdateOn;default:null;comment:Updated datetime(UTC)"`
	DeleteOn sql.NullTime `gorm:"column:DeleteOn;default:null;comment:Deleted datetime(UTC)"`
}

type Todo struct {
	Id     uuid.UUID `gorm:"primarykey;type:uuid;column:Id;default:uuid_generate_v4()"`
	Title  string    `gorm:"column:Title;not null"`
	IsDone bool      `gorm:"column:IsDone;not null;default:false"`
	TrackDateTimes
	// gorm.Model // We can embeded the gorm.Model that has CreatedAt, UpdatedAt and DeletedAt fields
}

type TodoDetail struct {
	Id          uuid.UUID `gorm:"primaryKey;type:uuid;column:Id"`
	Priority    int       `gorm:"column:Priority;"`
	Description string    `gorm:"column:Description;size:500"`
}

type Priority struct {
	Id   int    `gorm:"primaryKey;autoIncrement:true;"`
	Name string `gorm:"unique;column:Name;size:20;not null"`
}

// TableName: Specified table name for struct Todo
func (Todo) TableName() string {
	return "Todos"
}

func (TodoDetail) TableName() string {
	return "TodoDetails"
}

func (Priority) TableName() string {
	return "Priorities"
}

func (m *Todo) Print() {
	om, _ := json.MarshalIndent(m, "", "\t")
	fmt.Printf("%s\n", string(om))
}
