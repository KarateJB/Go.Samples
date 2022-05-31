package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"example/graphql/graph/generated"
	"example/graphql/graph/model"
	"fmt"

	"github.com/google/uuid"
	"github.com/stroiman/go-automapper"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	var user *model.User
	automapper.MapLoose(input, &user)
	r.users = append(r.users, user)
	return user, nil
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	// var user *model.User
	// if input.UserId != "" {
	// 	for _, u := range r.users {
	// 		if u.Id == input.UserId {
	// 			user = u
	// 			break
	// 		}
	// 	}
	// }

	todo := &model.Todo{
		Id:     uuid.New(),
		Title:  input.Title,
		IsDone: input.IsDone,
		UserId: input.UserId,
	}

	r.todos = append(r.todos, todo)
	return todo, nil
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
	panic(fmt.Errorf("not implemented"))
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	for _, user := range r.users {
		if user.Id == obj.UserId {
			return user, nil
		}
	}

	return nil, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
