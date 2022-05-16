package main

import (
	"log"
	"types"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=1qaz2wsx dbname=postgres port=5432 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database")
	}

	// Migrate
	db.AutoMigrate(&types.Todo{})

	// Initialize data
	initData(db)
}

// initData: Initialize data
func initData(db *gorm.DB) {
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
