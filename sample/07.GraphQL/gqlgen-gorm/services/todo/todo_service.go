package todoservice

import (
	"database/sql"
	models "example/graphql/graph/model"
	dbtypes "example/graphql/types/db"
	"example/graphql/utils"
	"fmt"
	"time"

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

// GetAll: get all TODOs
func (m *TodoAccess) GetAll() []*models.Todo {
	var entities []dbtypes.Todo
	var todos []*models.Todo
	var count int64

	m.DB.Preload(clause.Associations).Preload("TodoExt.Priority").Find(&entities).Count(&count)

	if count > 0 {
		for _, entity := range entities {
			var todo *models.Todo
			automapper.MapLoose(entity, &todo)
			todos = append(todos, todo)
		}
		return todos
	} else {
		return nil
	}
}

// GetOne: get the TODO by Id
func (m *TodoAccess) GetOne(id uuid.UUID) *models.Todo {
	var entity *dbtypes.Todo
	var todo *models.Todo
	var count int64

	m.DB.Model(&dbtypes.Todo{}).Where(`"Id" = ?`, id).Preload(clause.Associations).Preload("TodoExt.Priority").First(&entity).Count(&count)
	// m.DB.Model(&dbtypes.Todo{}).Where(`"Id" = ?`, id).Joins("TodoExt").First(&entity).Count(&count) // Or specify the outer join target

	if count > 0 {
		automapper.MapLoose(entity, &todo)
	}
	return todo
}

// Search: search TODOs that its Title contains "queryValTitle" and IsDone matchs "queryValIsDone"
func (m *TodoAccess) Search(queryValTitle string, queryValIsDone bool) *[]models.Todo {
	var entities *[]dbtypes.Todo
	var todos []models.Todo
	var count int64

	m.DB.Model(entities).
		Where(`"Title" LIKE ?`, fmt.Sprintf("%%%s%%", queryValTitle)). // Or use strings.Builder
		Where(`"IsDone" = ?`, queryValIsDone).Preload(clause.Associations).
		Preload("TodoExt.Priority").Find(&entities).Count(&count)

	if count > 0 {
		for _, entity := range *entities {
			var todo models.Todo
			automapper.MapLoose(entity, &todo)
			todos = append(todos, todo)
		}
		return &todos
	} else {
		return nil
	}

}

// Create: create a new TODO
func (m *TodoAccess) Create(todo *models.NewTodo) *models.Todo {
	var entity dbtypes.Todo
	automapper.MapLoose(todo, &entity)
	m.DB.Create(&entity)
	m.DB.Model(&entity).Association("Tags").Append(entity.Tags)

	// Optional: if we use the custom many-to-many relation table "TodoTags"
	todoTags := utils.Map(todo.TagIds, func(tagId uuid.UUID) dbtypes.TodoTag {
		return dbtypes.TodoTag{TodoId: entity.Id, TagId: tagId}
	})
	m.DB.Create(&todoTags)

	createdTodo := m.GetOne(entity.Id)
	return createdTodo
}

// Update: update the TODO
func (m *TodoAccess) Update(todo *models.EditTodo) (*models.Todo, int64) {
	var entity dbtypes.Todo
	var updatedCount int64

	automapper.MapLoose(todo, &entity)

	// Update TODO
	entity.TrackDateTimes = dbtypes.TrackDateTimes{
		UpdateOn: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}
	m.DB.Model(&dbtypes.Todo{}).Where(`"Id" = ?`, todo.Id).Updates(&entity).Count(&updatedCount)

	// Update TodoExt
	// m.DB.Model(&entity).Association("TodoExt").Append(&entity.TodoExt) // Not work, see https://github.com/go-gorm/gorm/issues/3487
	// m.DB.Session(&gorm.Session{FullSaveAssociations: true}).Save(&entity) // This will work
	m.DB.Model(dbtypes.TodoExt{}).Where(`"Id" = ?`, todo.TodoExt.Id).Updates(&entity.TodoExt)

	// Update todo_tages
	m.DB.Model(&entity).Association("Tags").Replace(entity.Tags)

	// Optional: if we use the custom many-to-many relation table "TodoTags"
	m.DB.Model(&dbtypes.TodoTag{}).Where(`"TodoId" = ?`, todo.Id).Delete(&dbtypes.TodoTag{})
	todoTags := utils.Map(todo.TagIds, func(tagId uuid.UUID) dbtypes.TodoTag {
		return dbtypes.TodoTag{TodoId: todo.Id, TagId: tagId}
	})
	m.DB.Create(&todoTags)

	// Query the latest TODO
	updatedTodo := m.GetOne(todo.Id)

	return updatedTodo, updatedCount
}

// DeleteOne: delete a TODO
func (m *TodoAccess) DeleteOne(id uuid.UUID) int64 {
	// Optional: If we don't set CASCADE delete, then we can use Association to remove the relations, e.q. relations in todo_tags as following.
	// var entity dbtypes.Todo
	// m.DB.Model(&dbtypes.Todo{}).Where(`"Id" = ?`, id).Preload("Tags").First(&entity)
	// m.DB.Model(&entity).Association("Tags").Delete(entity.Tags)

	var count int64
	m.DB.Model(&dbtypes.Todo{}).Where(`"Id" = ?`, id).Count(&count).Delete(&dbtypes.Todo{})

	return count
}

// Delete: delete one or more TODOs
func (m *TodoAccess) Delete(todoIds *[]uuid.UUID) int64 {
	var count int64
	var entities []dbtypes.Todo

	// m.DB.Model(&dbtypes.Todo{}).Where()
	m.DB.Find(&entities, todoIds).Count(&count).Delete(&dbtypes.Todo{})
	// db.Delete(&types.Todo{}, `"Id" IN ?`, todoIds)

	return count
}
