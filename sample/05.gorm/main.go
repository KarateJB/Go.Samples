package main

import (
	"database/sql"
	"log"
	"time"
	"types"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

const logLevel = logger.Info // logger.Info

func main() {
	// Set database connection string
	dsn := "host=localhost user=postgres password=1qaz2wsx dbname=postgres port=5432 sslmode=disable TimeZone=UTC"

	openedDb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logLevel)})
	if err != nil {
		log.Fatal("Failed to connect database")
	}
	db = openedDb

	// Migrate
	db.AutoMigrate(&types.Todo{})

	// Initialize data
	initData()

	// Single row handling
	// handleSingleRow()

	// Multiple rows handling
	handleMultipleRows()
}

// initData: Initialize data
func initData() {
	// Create
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
	// Add a TODO
	id := uuid.MustParse("aa3cdd2f-17b9-4f43-9eb0-af56b42908c5")
	todo := types.Todo{
		Id:     id,
		Title:  "Task A",
		IsDone: false,
	}
	db.FirstOrCreate(&todo) // Read the record that matchs the value of "id", or insert a new row.
	todo.Print()

	// Read a TODO
	var existTodo types.Todo
	db.First(&existTodo, id)
	existTodo.Print()

	// Update a TODO
	db.Model(&existTodo).Update("Title", "Task ???")
	db.Model(&existTodo).Updates(types.Todo{
		IsDone: true,
		TrackDateTimes: types.TrackDateTimes{
			UpdateOn: sql.NullTime{Time: time.Now(), Valid: true}, // Set Valid = true is optional, it will be true once we read the row from DB.
		},
	})
	existTodo.Print()

	// Delete a TODO
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
	db.Model(types.Todo{}).Where(`"IsDone" = ?`, false).Updates(types.Todo{
		IsDone: true,
		TrackDateTimes: types.TrackDateTimes{
			UpdateOn: sql.NullTime{Time: time.Now(), Valid: true},
		},
	})
}
