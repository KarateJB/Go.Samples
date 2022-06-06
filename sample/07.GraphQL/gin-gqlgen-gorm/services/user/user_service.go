package userservice

import (
	models "example/graphql/graph/model"
	"example/graphql/types"
	dbtypes "example/graphql/types/db"

	"github.com/stroiman/go-automapper"
	"gorm.io/gorm"
)

type UserAccess struct {
	DB   *gorm.DB
	Mode types.Mode
}

// New: UserService factory
func New(db *gorm.DB, mode types.Mode) *UserAccess {
	return &UserAccess{
		DB:   db,
		Mode: mode,
	}
}

// Get: get the user by Id
func (m *UserAccess) GetOne(id string) *models.User {
	var entity *dbtypes.User
	var user *models.User
	var cnt int64
	m.DB.First(&entity, `"Id" = ?`, id).Count(&cnt)
	// m.DB.Model(&dbtypes.User{}).Where(`"Id" = ?`, id).First(&entity)
	if cnt > 0 {
		automapper.MapLoose(entity, &user)
	}
	return user
}

// GetAll: get all users
func (m *UserAccess) GetAll() []*models.User {
	var entities []dbtypes.User
	var users []*models.User
	var cnt int64

	m.DB.Find(&entities).Count(&cnt)

	if cnt > 0 {
		for _, entity := range entities {
			var user *models.User
			automapper.MapLoose(entity, &user)
			users = append(users, user)
		}
		return users
	} else {
		return nil
	}
}

// Create: create a new user
func (m *UserAccess) Create(user *models.NewUser) *models.User {
	var entity dbtypes.User
	var createdUser *models.User
	automapper.MapLoose(user, &entity)
	m.DB.Create(&entity)

	automapper.MapLoose(entity, &createdUser)
	return createdUser
}

// Update: update a user
func (m *UserAccess) Update(user *models.EditUser) (*models.User, int64) {
	var entity *dbtypes.User
	var updatedCnt int64
	var updatedUser *models.User
	m.DB.Model(&dbtypes.User{}).Where(`"Id" = ?`, user.Id).Updates(dbtypes.User{
		Name: user.Name,
	}).First(&entity).Count(&updatedCnt)

	automapper.MapLoose(entity, &updatedUser)
	return updatedUser, updatedCnt
}

// Delete: delete a user
func (m *UserAccess) Delete(id string) int64 {
	var cnt int64
	m.DB.Model(&dbtypes.User{}).Where(`"Id" = ?`, id).Count(&cnt).Delete(&dbtypes.User{})
	return cnt
}
