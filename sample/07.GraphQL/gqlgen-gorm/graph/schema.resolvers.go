package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"example/graphql/graph/generated"
	"example/graphql/graph/model"
	dbservice "example/graphql/services/db"
	todoservice "example/graphql/services/todo"
	userservice "example/graphql/services/user"
	"fmt"

	"github.com/google/uuid"
)

var (
	userService *userservice.UserAccess = userservice.New((dbservice.New()).DB)
	todoService *todoservice.TodoAccess = todoservice.New((dbservice.New()).DB)
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	createduser := userService.Create(&input)
	return createduser, nil
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	createdTodo := todoService.Create(&input)
	return createdTodo, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.EditUser) (*model.User, error) {
	updatedUser, updatedCount := userService.Update(&input)
	if updatedCount == 1 {
		return updatedUser, nil
	} else {
		return nil, fmt.Errorf("User (Id: %s) not found", input.Id)
	}
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, input model.EditTodo) (*model.Todo, error) {
	updatedTodo, updatedCount := todoService.Update(&input)
	if updatedCount == 1 {
		return updatedTodo, nil
	} else {
		return nil, fmt.Errorf("TODO (Id: %s) not found", input.Id.String())
	}
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	const deleteOk = true
	if cnt := userService.Delete(id); cnt == 1 {
		return deleteOk, nil
	} else {
		return !deleteOk, fmt.Errorf("User (Id: %s) not found", id)
	}
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, id uuid.UUID) (bool, error) {
	const deleteOk = true
	if cnt := todoService.DeleteOne(id); cnt == 1 {
		return deleteOk, nil
	} else {
		return !deleteOk, fmt.Errorf("TODO (Id: %s) not found", id.String())
	}
}

func (r *mutationResolver) DeleteTodos(ctx context.Context, input []uuid.UUID) (*int64, error) {
	deletedCnt := todoService.Delete(&input)
	return &deletedCnt, nil
}

func (r *queryResolver) Todo(ctx context.Context, id uuid.UUID) (*model.Todo, error) {
	todo := todoService.GetOne(id)
	return todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	todos := todoService.GetAll()
	return todos, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	user := userService.GetOne(id)
	return user, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	users := userService.GetAll()
	return users, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
