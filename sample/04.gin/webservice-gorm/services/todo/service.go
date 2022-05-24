package todoservice

import (
	types "example/webservice/types/api"
	dbtypes "example/webservice/types/db"

	"github.com/google/uuid"
	"github.com/stroiman/go-automapper"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TodoAccess struct {
	DB *gorm.DB
}

// New: TodoService factory
func New(db *gorm.DB) *TodoAccess {
	return &TodoAccess{
		DB: db,
	}
}

// Get: get the todo by Id
func (m *TodoAccess) Get(id uuid.UUID) *types.Todo {
	var entity *dbtypes.Todo
	var todo *types.Todo
	var count int64

	m.DB.Model(&dbtypes.Todo{}).Where(`"Id" = ?`, id).Preload(clause.Associations).Preload("TodoExt.Priority").First(&entity).Count(&count)
	// m.DB.Model(&dbtypes.Todo{}).Where(`"Id" = ?`, id).Joins("TodoExt").First(&entity).Count(&count)
	// m.DB.Model(&dbtypes.Todo{}).Where(`"Id" = ?`, id).Preload(clause.Associations).First(&entity).Count(&count)

	if count > 0 {
		automapper.MapLoose(entity, &todo)
	}
	return todo
}

// Create: create a new todo
func (m *TodoAccess) Create(todo *types.Todo) *dbtypes.Todo {
	var entity dbtypes.Todo
	automapper.MapLoose(todo, &entity)
	m.DB.Create(&entity)
	m.DB.Model(&entity).Association("Tags").Append(entity.Tags)
	return &entity
}

// Update: update a todo
func (m *TodoAccess) Update(todo *types.Todo) int64 {
	var updatedCount int64

	// Update TODO
	m.DB.Model(&dbtypes.Todo{}).Where(`"Id" = ?`, todo.Id).Updates(dbtypes.Todo{
		Title:  todo.Title,
		IsDone: todo.IsDone,
		TodoExt: dbtypes.TodoExt{
			Description: todo.TodoExt.Description,
			PriorityId:  todo.TodoExt.PriorityId,
		},
		UserId: todo.UserId,
	}).Count(&updatedCount)

	// Update (Delete and create) tags mapping
	m.DB.Model(&dbtypes.TodoTag{}).Where(`"TodoId" = ?`, todo.Id).Delete(&dbtypes.TodoTag{})
	tagIds := Map(todo.Tags, func(tag types.Tag) uuid.UUID {
		return tag.Id
	})
	todoTags := Map(tagIds, func(tagId uuid.UUID) dbtypes.TodoTag {
		return dbtypes.TodoTag{TodoId: todo.Id, TagId: tagId}
	})
	m.DB.Create(&todoTags)

	return updatedCount
}

// Delete: delete a todo
func (m *TodoAccess) Delete(todo *types.Todo) int64 {
	var count int64
	m.DB.Model(&dbtypes.Todo{}).Where(`"Id" = ?`, todo.Id).Count(&count).Delete(&dbtypes.Todo{})
	return count
}

// Map: map A array to B array
// TODO: refactor
func Map[S, D any](src []S, f func(S) D) []D {
	us := make([]D, len(src))
	for i := range src {
		us[i] = f(src[i])
	}
	return us
}

func getTagId(src types.Tag) uuid.UUID {
	return src.Id
}
