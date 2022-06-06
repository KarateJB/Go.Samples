package main

import (
	todoapi "example/graphql/api/todo"
	userapi "example/graphql/api/user"
	"example/graphql/config"
	"example/graphql/graph"
	"example/graphql/graph/generated"
	dbservice "example/graphql/services/db"
	"fmt"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func main() {

	/* Database */
	initDb()

	// Gin
	router := gin.Default()

	// GraphQL
	router.POST("/query", graphqlHandler())
	router.GET("/", playgroundHandler())

	// RESTful
	apiRouterGroup := router.Group("/api")
	{
		userRg := apiRouterGroup.Group("/user")
		{
			userRg.GET(":id", userapi.GetUser)
			userRg.POST("", userapi.PostUser)
			userRg.PUT("", userapi.PutUser)
			userRg.DELETE("", userapi.DeleteUser)
		}
		todoRg := apiRouterGroup.Group("/todo")
		{
			todoRg.GET(":id", todoapi.GetTodo) // The id is required for matching this routing
			// todoRg.GET("*id", todoapi.GetTodoById) // The id is optional for matching this routing, e.q. api/todo/ or api/todo/xxx
			todoRg.POST("", todoapi.PostTodo)
			todoRg.PUT("", todoapi.PutTodo)
			todoRg.DELETE("", todoapi.DeleteTodo)
		}
		todosRg := apiRouterGroup.Group("/todos")
		{
			todosRg.GET("", todoapi.GetAllTodos)
			todosRg.GET("search", todoapi.SearchTodo)
			todosRg.DELETE("", todoapi.DeleteTodos)
		}
	}

	// port := os.Getenv("GQLGEN_PORT") // From env var
	configs := config.Init()
	router.Run(fmt.Sprintf("localhost:%s", configs.Port))
}

// initDb: Initialize database
func initDb() {
	// DB connect configuration
	dbService := dbservice.New()
	dbService.Migrate()
	dbService.InitData()
}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	// h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	gqlConfig := generated.Config{Resolvers: &graph.Resolver{}}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(gqlConfig))
	return func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
