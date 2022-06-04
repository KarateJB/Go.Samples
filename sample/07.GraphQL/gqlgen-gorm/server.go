package main

import (
	"example/graphql/config"
	"example/graphql/graph"
	directives "example/graphql/graph/directive"
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
	gqlConfig := generated.Config{Resolvers: &graph.Resolver{}}
	gqlConfig.Directives.MaskUserName = directives.MaskUserName
	gqlConfig.Directives.HasTag = directives.HasTag
	gqlConfig.Directives.CheckRules = directives.CheckRules
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
