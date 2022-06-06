package dbservice

import (
	config "example/graphql/config"
	dbtypes "example/graphql/types/db"
	"log"
	"strconv"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbAccess struct {
	DB    *gorm.DB
	Error error
}

const LogLevel logger.LogLevel = logger.Info

// New: create and get the Database access instance
func New() *DbAccess {
	// Get configuration
	configs := config.Init()

	dsn := configs.DB
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(LogLevel)})
	if err != nil {
		log.Fatal("Failed to connect database")
	}

	return &DbAccess{
		DB: db,
	}
}

// Migrate: database migration
func (m *DbAccess) Migrate() {
	// Migrate
	m.DB.AutoMigrate(&dbtypes.Priority{}, &dbtypes.User{}, &dbtypes.Todo{}, &dbtypes.TodoExt{}, &dbtypes.Tag{}, &dbtypes.TodoTag{})
}

// InitData: Initialize data
func (m *DbAccess) InitData() {
	// Init Priorities
	m.DB.Create(&[]dbtypes.Priority{
		{Id: 1, Name: "High"},
		{Id: 2, Name: "Medium"},
		{Id: 3, Name: "Low"},
	})

	// Test data
	m.DB.Create(&dbtypes.User{
		Id:    "JB_" + strconv.FormatInt(time.Now().Unix(), 10),
		Name:  "JB Lin",
		Todos: *dbtypes.Todo{}.CreateRandom(3),
	})

	m.DB.Create(&[]dbtypes.Tag{
		{Id: uuid.MustParse("6aee5542-3f70-4cbc-ab05-fd020285f021"), Name: "DevOps"},
		{Id: uuid.MustParse("dcc5a568-ae07-4600-9055-97eb129f319c"), Name: "Programming"},
	})
}
