package directives

import (
	"context"
	models "example/graphql/graph/model"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"golang.org/x/exp/slices"
)

// MaskUserName: mask user's name
func MaskUserName(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	user := obj.(*models.User)
	if sz := len(user.Name); sz > 1 {
		user.Name = fmt.Sprint(user.Name[:sz-(sz/2)] + "***")
	}

	return next(ctx) // let it pass through
	// return nil, fmt.Errorf("error") // or block calling the next resolver
}

// WIP
// HasTag: make sure that a TODO has at least one tag
func HasTag(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	// val, err := next(ctx) // HACK: obj is nil, see https://github.com/99designs/gqlgen/issues/1084
	todo := obj.(*models.NewTodo)
	if todo.Tags == nil || len(todo.Tags) == 0 {
		return nil, fmt.Errorf("a TODO must have at least one tag") // Block calling the next resolver
	} else {
		return next(ctx)
		// return val, err
	}
}

// WIP
// CheckRules: check rules for TODO
func CheckRules(ctx context.Context, obj interface{}, next graphql.Resolver, rules []models.Rule) (interface{}, error) {
	// val, err := next(ctx) // HACK: obj is nil, see https://github.com/99designs/gqlgen/issues/1084
	const MAX_TAGS int = 2
	todo := obj.(*models.NewTodo)

	// Start checking rules
	if slices.Contains(rules, models.RuleLimitedTag) && len(todo.Tags) > MAX_TAGS {
		return nil, fmt.Errorf("the TODO has %d tags, that it is allowed to have at most %d tags", len(todo.Tags), MAX_TAGS)
	}
	if slices.Contains(rules, models.RuleHasOwner) && todo.UserId == "" {
		return nil, fmt.Errorf("the TODO must have an owner (field: `user`)")
	}

	return next(ctx)
	// return val, err
}
