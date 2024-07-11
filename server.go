package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/sheeiavellie/ozon050724/graph"
	"github.com/sheeiavellie/ozon050724/storage"
)

const defaultPort = "8080"
const defaultHost = "localhost"

func main() {
	// env stuff
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	host := os.Getenv("HOST")
	if host == "" {
		host = defaultHost
	}

	// str selection
	var str storage.Storage

	postgresStorage, err := storage.NewPostgresStorage()
	if err != nil {
		return
	}
	str = postgresStorage

	// server set up
	resolver := graph.NewResolver(str)

	srv := handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{Resolvers: resolver},
		),
	)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://%s:%s/ for GraphQL playground", host, port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
