package userservice

import (
	types "example/webservice/types/api"
	dbtypes "example/webservice/types/db"

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

// Get: get the user by Id
func (m *UserAccess) Get(id string) *types.User {
	var entity *dbtypes.User
	var user *types.User
	var count int64
	m.DB.First(&entity, `"Id" = ?`, id).Count(&count)
	// m.DB.Model(&dbtypes.User{}).Where(`"Id" = ?`, id).First(&entity)
	if count > 0 {
		automapper.MapLoose(entity, &user)
	}
	return user
}

// Create: create a new user
func (m *UserAccess) Create(user *types.User) {
	var entity dbtypes.User
	automapper.MapLoose(user, &entity)
	m.DB.Create(&entity)
}

// Update: update a user
func (m *UserAccess) Update(user *types.User) {
	var entity dbtypes.User
	m.DB.First(&entity, user.Id).Updates(dbtypes.User{
		Name: user.Name,
	})
}

// Delete: delete a user
func (m *UserAccess) Delete(user *types.User) int64 {
	var count int64
	m.DB.Model(&dbtypes.User{}).Where(`"Id" = ?`, user.Id).Count(&count).Delete(&dbtypes.User{})
	return count
}
