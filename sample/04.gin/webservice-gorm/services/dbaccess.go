package dbaccess

import (
	"log"
	"types"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbAccess struct {
	Database *gorm.DB
	Error    error
}

func New(dsn string, logLevel logger.LogLevel) *DbAccess {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logLevel)})
	if err != nil {
		log.Fatal("Failed to connect database")
	}

	return &DbAccess{
		Database: db,
	}
}

func (m *DbAccess) Create(entity types.User) {

}
