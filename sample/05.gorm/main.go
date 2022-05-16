package main

import (
	"database/sql"
	"log"
	"time"
	"types"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	dsn := "host=localhost user=postgres password=1qaz2wsx dbname=postgres port=5432 sslmode=disable TimeZone=UTC"
	openedDb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database")
	}
	db = openedDb

	// Migrate
	db.AutoMigrate(&types.Todo{})

	// Initialize data
	initData()

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
	// db.Create(&types.Todo{
	// 	Id:     uuid.MustParse("aa3cdd2f-17b9-4f43-9eb0-af56b42908c5"),
	// 	Title:  "Task A",
	// 	IsDone: false})
}
