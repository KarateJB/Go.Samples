package dbaccess

import (
	dbtypes "example/webservice/types/data_access"
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

// New: create and get the Database access instance
func New(dsn string, logLevel logger.LogLevel) *DbAccess {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logLevel)})
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
		{Id: uuid.New(), Name: "DevOps"},
		{Id: uuid.New(), Name: "Programming"},
	})
}

func (m *DbAccess) Create(entity interface{}) {
	_ = m.DB.Create(entity)
}

func (m *DbAccess) BatchCreate(entities []interface{}, batchSize *int) {
	if batchSize == nil {
		m.DB.Create(entities)
	} else {
		m.DB.CreateInBatches(entities, *batchSize)
	}
}
