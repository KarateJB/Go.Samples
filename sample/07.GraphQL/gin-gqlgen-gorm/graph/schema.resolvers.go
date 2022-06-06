package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"example/graphql/graph/generated"
	"example/graphql/graph/model"
	"example/graphql/services"
	"fmt"

	"github.com/google/uuid"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	createduser := services.UserGql.Create(&input)
	return createduser, nil
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	createdTodo := services.TodoGql.Create(&input)
	return createdTodo, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.EditUser) (*model.User, error) {
	updatedUser, updatedCount := services.UserGql.Update(&input)
	if updatedCount == 1 {
		return updatedUser, nil
	} else {
		return nil, fmt.Errorf("User (Id: %s) not found or failed to update", input.Id)
	}
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, input model.EditTodo) (*model.Todo, error) {
	updatedTodo, updatedCount := services.TodoGql.Update(&input)
	if updatedCount == 1 {
		return updatedTodo, nil
	} else {
		return nil, fmt.Errorf("TODO (Id: %s) not found or failed to update", input.Id.String())
	}
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	const deleteOk = true
	if cnt := services.UserGql.Delete(id); cnt == 1 {
		return deleteOk, nil
	} else {
		return !deleteOk, fmt.Errorf("User (Id: %s) not found or failed to delete", id)
	}
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, id uuid.UUID) (bool, error) {
	const deleteOk = true
	if cnt := services.TodoGql.DeleteOne(id); cnt == 1 {
		return deleteOk, nil
	} else {
		return !deleteOk, fmt.Errorf("TODO (Id: %s) not found or failed to delete", id.String())
	}
}

func (r *mutationResolver) DeleteTodos(ctx context.Context, input []uuid.UUID) (*int64, error) {
	deletedCnt := services.TodoGql.Delete(&input)
	return &deletedCnt, nil
}

func (r *queryResolver) Todo(ctx context.Context, id uuid.UUID) (*model.Todo, error) {
	todo := services.TodoGql.GetOne(id)
	return todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	todos := services.TodoGql.GetAll()
	return todos, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	user := services.UserGql.GetOne(id)
	return user, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	users := services.UserGql.GetAll()
	return users, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
