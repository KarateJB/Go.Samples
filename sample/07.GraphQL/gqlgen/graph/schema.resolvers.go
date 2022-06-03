package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"example/graphql/graph/generated"
	"example/graphql/graph/model"

	"github.com/google/uuid"
	"github.com/stroiman/go-automapper"
	"golang.org/x/exp/slices"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	var user *model.User
	automapper.MapLoose(input, &user)
	r.users = append(r.users, user)
	return user, nil
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Id:     uuid.New(),
		Title:  input.Title,
		IsDone: input.IsDone,
		UserId: input.UserId,
	}

	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.EditUser) (*model.User, error) {
	for index, user := range r.users {
		if user.Id == input.Id {
			r.users[index].Name = input.Name
			return r.users[index], nil
		}
	}

	return nil, nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, input model.EditTodo) (*model.Todo, error) {
	for index, todo := range r.todos {
		if todo.Id == input.Id {
			r.todos[index].Title, r.todos[index].IsDone = input.Title, input.IsDone
			return r.todos[index], nil
		}
	}

	return nil, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	const deleteOk = true
	for index, user := range r.users {
		if user.Id == id {
			r.users = slices.Delete(r.users, index, index+1)
			return deleteOk, nil
		}
	}

	return !deleteOk, nil
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, id uuid.UUID) (bool, error) {
	const deleteOk = true
	for index, todo := range r.todos {
		if todo.Id == id {
			r.todos = slices.Delete(r.todos, index, index+1)
			return deleteOk, nil
		}
	}

	return !deleteOk, nil
}

func (r *queryResolver) Todo(ctx context.Context, id string) (*model.Todo, error) {
	for _, todo := range r.todos {
		if todo.Id.String() == id {
			return todo, nil
		}
	}

	return nil, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	for _, user := range r.users {
		if user.Id == id {
			return user, nil
		}
	}

	return nil, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return r.users, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
