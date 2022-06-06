package services

import (
	dbservice "example/graphql/services/db"
	todoservice "example/graphql/services/todo"
	userservice "example/graphql/services/user"
	"example/graphql/types"
)

var (
	UserRf  *userservice.UserAccess = userservice.New((dbservice.New()).DB, types.RestFul)
	TodoRf  *todoservice.TodoAccess = todoservice.New((dbservice.New()).DB, types.RestFul)
	UserGql *userservice.UserAccess = userservice.New((dbservice.New()).DB, types.GraphQL)
	TodoGql *todoservice.TodoAccess = todoservice.New((dbservice.New()).DB, types.GraphQL)
)
