package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"example/graphql/graph/generated"
	"example/graphql/graph/model"
	"fmt"
)

func (r *todoResolver) TodoExt(ctx context.Context, obj *model.Todo) (*model.TodoExt, error) {
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

func (r *todoResolver) Tags(ctx context.Context, obj *model.Todo) ([]*model.Tag, error) {
	panic(fmt.Errorf("not implemented"))
}

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type todoResolver struct{ *Resolver }
