package services

import (
	dbservice "example/graphql/services/db"
	todoservice "example/graphql/services/todo"
	userservice "example/graphql/services/user"
)

var (
	User *userservice.UserAccess = userservice.New((dbservice.New()).DB)
	Todo *todoservice.TodoAccess = todoservice.New((dbservice.New()).DB)
)
