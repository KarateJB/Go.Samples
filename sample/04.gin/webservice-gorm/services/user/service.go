package userservice

import (
	types "example/webservice/types/api"
	dbtypes "example/webservice/types/data_access"

	"github.com/stroiman/go-automapper"
	"gorm.io/gorm"
)

type UserAccess struct {
	DB *gorm.DB
}

// New: UserService factory
func New(db *gorm.DB) *UserAccess {
	return &UserAccess{
		DB: db,
	}
}

// Create:
func (m *UserAccess) Create(user *types.User) {
	var entity dbtypes.User
	automapper.MapLoose(user, &entity)
	m.DB.Create(&entity)
}
