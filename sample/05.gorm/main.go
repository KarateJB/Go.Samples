package main

import (
	"fmt"
	"log"
	"time"
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

	shit := types.Todo{
		Id:     uuid.New(),
		Title:  "Test",
		IsDone: true,
		Model: gorm.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	fmt.Println(shit)

	db.AutoMigrate(&types.Todo{})

	// Create
	// db.Create(&shit)
	db.Create(&types.Todo{
		Id:     uuid.MustParse("aa3cdd2f-17b9-4f43-9eb0-af56b42908c5"),
		Title:  "Task A",
		IsDone: false})
}
