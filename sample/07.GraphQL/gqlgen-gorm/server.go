package main

import (
	"example/graphql/config"
	"example/graphql/graph"
	"example/graphql/graph/generated"
	dbservice "example/graphql/services/db"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {

	/* Database */
	initDb()

	/* GraphQL */
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

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
