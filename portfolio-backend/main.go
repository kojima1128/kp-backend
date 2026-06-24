package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kojima1128/portfolio-backend/internal/graph"
	"github.com/kojima1128/portfolio-backend/internal/infrastructure"
	mysqlrepo "github.com/kojima1128/portfolio-backend/internal/infrastructure/mysql"
	"github.com/kojima1128/portfolio-backend/internal/service"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db := infrastructure.NewDB()
	userRepo := mysqlrepo.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	resolver := graph.NewResolver(userService)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: resolver,
	}))

	http.Handle("/query", srv)
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
