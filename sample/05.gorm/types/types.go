package types

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type TrackDateTimes struct {
	CreateOn time.Time    `gorm:"column:CreateOn;not null;default:now();comment:Created datetime(UTC)"`
	UpdateOn sql.NullTime `gorm:"column:UpdateOn;default:NULL;comment:Updated datetime(UTC)"`
	DeleteOn sql.NullTime `gorm:"column:DeleteOn;default:NULL;comment:Deleted datetime(UTC)"`
}

// Owner: A owner has one or many Todo
type User struct {
	Id    string `gorm:"column:UserId;primaryKey;size:100"`
	Name  string `gorm:"column:Name;size:200;not null"`
	Todos []Todo `gorm:"foreignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // HasMany relation on Todo. "Todo" will has a foreign key "UserId" which has reference on "User"."Id"
}

// Todo: A Todo has only one TodoExt
type Todo struct {
	Id             uuid.UUID `gorm:"column:Id;type:uuid;primarykey;default:uuid_generate_v4()"`
	Title          string    `gorm:"column:Title;not null"`
	IsDone         bool      `gorm:"column:IsDone;not null;default:false"`
	TrackDateTimes           // Or use gorm.Model instead
	TodoExt        TodoExt   `gorm:"foreignkey:Id;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // HasOne relation on "TodoExt". The "TodoExt" will has a foreign key "Id" which has reference on "Todo"."Id"
	UserId         string    `gorm:"column:UserId;default:NULL"`                                                // If this field name is "UserId", we can ignore setting "foreignKey:Id" on "User" struct's field "Todos"
}

// TodoExt: Todo's extension table
type TodoExt struct {
	Id          uuid.UUID `gorm:"column:Id;primaryKey"` // HasOne relation on Todo. If this field name is "UserId", we can ignore setting "foreignKey:Id" on "Todo" struct's field "TodoExt"
	Description string    `gorm:"column:Description;size:500"`
	PriorityId  int       `gorm:"column:PriorityId"`
	Priority    Priority  `gorm:"foreignKey:PriorityId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // The tag: "foreignKey" is optional here, it uses the type name plus its primary field name in default.
}

// Priority: Mapping table
type Priority struct {
	Id   int    `gorm:"column:Id;primaryKey;autoIncrement:true;"`
	Name string `gorm:"column:Name;unique;size:20;not null"`
}

/* TableName: Specified table name for struct */
func (User) TableName() string {
	return "Users"
}

func (Todo) TableName() string {
	return "Todos"
}

func (TodoExt) TableName() string {
	return "TodoExts"
}

func (Priority) TableName() string {
	return "Priorities"
}

//------------------------------------------------

// Print: Output TODO as JSON string
func (m *Todo) Print() {
	om, _ := json.MarshalIndent(m, "", "\t")
	fmt.Printf("%s\n", string(om))
}

// CreateRand: Create a random TODO
func (t Todo) CreateRandom(n int) *[]Todo {
	rand.Seed(time.Now().UnixNano())
	var todos []Todo
	for i := 0; i < n; i++ {
		todos = append(todos, Todo{
			Id:     uuid.New(),
			Title:  "Random task",
			IsDone: false,
			TodoExt: TodoExt{
				PriorityId:  rand.Intn(3) + 1, //Random [1,3]
				Description: "Only for testing",
			},
			// Model: gorm.Model{
			// 	CreatedAt: time.Now(),
			// 	UpdatedAt: time.Now(),
			// },
		})
	}

	return &todos
}
