package directives

import (
	"context"
	models "example/graphql/graph/model"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
)

func MaskUserName(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {

	user := obj.(*models.User)
	if sz := len(user.Name); sz > 1 {
		user.Name = fmt.Sprint(user.Name[:sz-(sz/2)] + "***")
	}

	return next(ctx) // let it pass through
	// return nil, fmt.Errorf("error") // or block calling the next resolver
}
