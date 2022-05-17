package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"
	"types"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

const logLevel = logger.Silent // logger.Silent to disable outputing SQL tracking

func main() {
	// Set database connection string
	dsn := "host=localhost user=postgres password=1qaz2wsx dbname=postgres port=5432 sslmode=disable TimeZone=UTC"

	openedDb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logLevel)})
	if err != nil {
		log.Fatal("Failed to connect database")
	}
	db = openedDb

	// Migrate
	db.AutoMigrate(&types.Todo{}, types.Priority{})

	// Initialize data
	initData()

	// Single row handling
	// handleSingleRow()

	// Multiple rows handling
	// handleMultipleRows()
}

// initData: Initialize data
func initData() {
	// Init Priorities
	db.Create(&[]types.Priority{
		{Id: 1, Name: "High"},
		{Id: 2, Name: "Medium"},
		{Id: 3, Name: "Low"},
	})
	// New TODO
	newTodo := types.Todo{
		Id:     uuid.New(),
		Title:  "Test",
		IsDone: true,
		// Model: gorm.Model{
		// 	CreatedAt: time.Now(),
		// 	UpdatedAt: time.Now(),
		// },
	}
	db.Create(&newTodo)
	// db.Create(&types.Todo{...})
}

// handleSingleRow: Insert, update and delete single row samples
func handleSingleRow() {
	/* CREATE */
	id := uuid.MustParse("aa3cdd2f-17b9-4f43-9eb0-af56b42908c5")
	todo := types.Todo{
		Id:     id,
		Title:  "Test task",
		IsDone: false,
	}
	db.FirstOrCreate(&todo) // Read the record that matchs the value of "id", or insert a new row.
<<<<<<< HEAD
	fmt.Print("Created TODO as following...")
	todo.Print()
=======
	// todo.Print()
	Print(todo)
>>>>>>> 421ba59 (Create more DAO)

	/* READ */
	var existTodo types.Todo
<<<<<<< HEAD
	db.First(&existTodo, id) // Get first row by primary key
	// db.First(&existTodo, `"Id" = ?`, id) // Get first row by where condition
	// db.Model(&types.Todo{}).Where(`"Id" = ?`, id).First(&existTodo) // Get first row by where condition
	// db.Model(&types.Todo{}).Where(`"Id" = ?`, id).Where(`"Title" = ?`, "Test task").First(&existTodo) // AND conditions
	// db.Model(&types.Todo{}).Where(types.Todo{Id: id, Title: "Test task"}).First(&existTodo) // AND conditions
	// db.Model(&types.Todo{}).Where(`"Id" = ?`, id).Or(`"Title" = ?`, "Test task").First(&existTodo) // OR conditions
	fmt.Print("Read TODO as following...")
	existTodo.Print()

	/* UPDATE */
	// Update a TODO (After read it)
	db.First(&existTodo, id).Update("Title", "Task XXX")
	// db.Model(&existTodo).Update("Title", "Task XXX")
	// db.Model(&existTodo).Updates(types.Todo{
	// 	IsDone: true,
	// 	TrackDateTimes: types.TrackDateTimes{
	// 		UpdateOn: sql.NullTime{Time: time.Now(), Valid: true}, // Set Valid = true is optional, it will be true once we read the row from DB.
	// 	},
	// })
	fmt.Print("Updated TODO as following...")
	existTodo.Print()
=======
	db.First(&existTodo, id)
	// existTodo.Print()
	Print(existTodo)

	// Update a TODO
	db.Model(&existTodo).Update("Title", "Task ???")
	db.Model(&existTodo).Updates(types.Todo{
		IsDone: true,
		TrackDateTimes: types.TrackDateTimes{
			UpdateOn: sql.NullTime{Time: time.Now(), Valid: true}, // Set Valid = true is optional, it will be true once we read the row from DB.
		},
	})
	// existTodo.Print()
	Print(existTodo)
>>>>>>> 421ba59 (Create more DAO)

	// Update a TODO (without read it first)
	// db.Model(&types.Todo{}).Where(`"Id" = ?`, id).Update("Title", "Task XXX")

	/* DELETE */
	db.Delete(&existTodo)
}

// handleMultipleRows: Insert, update and delete multiple rows sample
func handleMultipleRows() {

	// Batch insert
	todos := []types.Todo{
		{Id: uuid.MustParse("0f36f6bc-5a26-4bf6-9557-75e2c5c9f12c"), Title: "Task A", IsDone: false},
		{Id: uuid.MustParse("6c3e7544-2d9d-46c0-a05a-d3c0390634b6"), Title: "Task B", IsDone: false},
		{Id: uuid.MustParse("d9f13086-413f-4583-a08b-62e3d4c7102e"), Title: "Task C", IsDone: false},
		{Id: uuid.MustParse("1bc6acb8-596e-4fcb-8514-23f42277d4a6"), Title: "Task D", IsDone: false},
		{Id: uuid.MustParse("5a5eddea-f904-4257-9532-c96522a2c169"), Title: "Task E", IsDone: false},
	}
	db.Create(&todos)
	// db.CreateInBatches(&todos, 3)

	// Batch update
	db.Model(&types.Todo{}).Where(`"IsDone" = ?`, false).Updates(types.Todo{
		IsDone: true,
		TrackDateTimes: types.TrackDateTimes{
			UpdateOn: sql.NullTime{Time: time.Now(), Valid: true},
		},
	})

	// Batch delete
	db.Where(`"Title" LIKE ?`, "Task%").Delete(&types.Todo{})
	// db.Delete(&types.Todo{}, `"Title" LIKE ?`, "Task%")
}

func Print[T any](m T) {
	om, _ := json.MarshalIndent(m, "", "\t")
	fmt.Printf("%s\n", string(om))
}
