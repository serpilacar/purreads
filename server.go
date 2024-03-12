package main

import (
	"log"
	"net/http"
	"os"
	"purreads/database"
	"purreads/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"

	"github.com/newrelic/go-agent/v3/newrelic"
)

const defaultPort = "8080"

func main() {
	godotenv.Load()
	config := &database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}
	db, err := database.NewConnection(config)
	if err != nil {
		panic(err)
	}
	database.Migrate(db)
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		DB: db,
	}}))

	if len(os.Getenv("NEW_RELIC_LICENSE")) != 0 {
		// initialize new relic
		_, newRelicErr := newrelic.NewApplication(
			newrelic.ConfigAppName("purreads"),
			newrelic.ConfigLicense(os.Getenv("NEW_RELIC_LICENSE")),
			newrelic.ConfigAppLogForwardingEnabled(true),
		)
		if newRelicErr != nil {
			panic(newRelicErr)
		}
	}

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
