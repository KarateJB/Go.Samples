package main

import (
	"context"
	"example/graphql/config"
	"example/graphql/graph"
	"example/graphql/graph/generated"
	models "example/graphql/graph/model"
	dbservice "example/graphql/services/db"
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {

	/* Database */
	initDb()

	/* GraphQL */
	gqlConfig := generated.Config{Resolvers: &graph.Resolver{}}
	gqlConfig.Directives.Mask = func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {

		user := obj.(*models.User)
		if sz := len(user.Name); sz > 1 {
			user.Name = fmt.Sprint(user.Name[:sz-(sz/2)] + "***")
		}

		return next(ctx) // let it pass through
		// return nil, fmt.Errorf("error") // or block calling the next resolver
	}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(gqlConfig))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	// port := os.Getenv("GQLGEN_PORT") // From env var
	configs := config.Init()
	port := configs.Port
	log.Printf("Go to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe("localhost:"+port, nil))
}

// initDb: Initialize database
func initDb() {
	// DB connect configuration
	dbService := dbservice.New()
	dbService.Migrate()
	dbService.InitData()
}
