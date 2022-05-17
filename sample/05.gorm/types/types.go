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
	UpdateOn sql.NullTime `gorm:"column:UpdateOn;default:NULL;comment:Updated datetime(UTC)"`
	DeleteOn sql.NullTime `gorm:"column:DeleteOn;default:NULL;comment:Deleted datetime(UTC)"`
}

// Todo:
type Todo struct {
	Id        uuid.UUID `gorm:"primarykey;type:uuid;column:Id;default:uuid_generate_v4()"`
	Title     string    `gorm:"column:Title;not null"`
	IsDone    bool      `gorm:"column:IsDone;not null;default:false"`
	TodoExtId uuid.UUID `gorm:"column:TodoExtId;default:NULL"`
	TrackDateTimes
	TodoExt TodoExt `gorm:"foreignKey:TodoExtId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	// gorm.Model // We can embeded the gorm.Model that has CreatedAt, UpdatedAt and DeletedAt fields
}

// TodoExt: Todo's extension table
type TodoExt struct {
	Id          uuid.UUID `gorm:"primaryKey;type:uuid;column:Id"`
	Description string    `gorm:"column:Description;size:500"`
	PriorityId  int       `gorm:"column:PriorityId"`
	Priority    Priority  `gorm:"foreignKey:PriorityId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // The tag: "foreignKey" is optional here, it uses the type name plus its primary field name in default.
}

// Priority: Mapping table
type Priority struct {
	Id   int    `gorm:"primaryKey;autoIncrement:true;"`
	Name string `gorm:"unique;column:Name;size:20;not null"`
}

// TableName: Specified table name for struct Todo
func (Todo) TableName() string {
	return "Todos"
}

func (TodoExt) TableName() string {
	return "TodoExts"
}

func (Priority) TableName() string {
	return "Priorities"
}

func (m *Todo) Print() {
	om, _ := json.MarshalIndent(m, "", "\t")
	fmt.Printf("%s\n", string(om))
}
