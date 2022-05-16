package types

import (
	"time"

	"github.com/google/uuid"
)

type TrackDateTimes struct {
	CreateOn time.Time `gorm:"column:CreateOn;not null;default:now();comment:Created datetime(UTC)"`
	UpdateOn time.Time `gorm:"column:UpdateOn;default:null;comment:Updated datetime(UTC)"`
	DeleteOn time.Time `gorm:"column:DeleteOn;default:null;comment:Deleted datetime(UTC)"`
}

type Todo struct {
	Id     uuid.UUID `gorm:"primarykey;type:uuid;column:Id;default:uuid_generate_v4()"`
	Title  string    `gorm:"column:Title;not null"`
	IsDone bool      `gorm:"column:IsDone;not null;default:false"`
	TrackDateTimes
	// gorm.Model // We can embeded the gorm.Model that has CreatedAt, UpdatedAt and DeletedAt fields
}

// TableName: Specified table name for struct Todo
func (Todo) TableName() string {
	return "Todos"
}
