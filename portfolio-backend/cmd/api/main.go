package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kojima1128/portfolio-backend/internal/common"
	"github.com/kojima1128/portfolio-backend/internal/graph"
	"github.com/kojima1128/portfolio-backend/internal/infrastructure"
	mysqlrepo "github.com/kojima1128/portfolio-backend/internal/infrastructure/mysql"
	"github.com/kojima1128/portfolio-backend/internal/service"
)

const defaultPort = "8080"

func main() {
	port := common.GetEnv("PORT", defaultPort)

	db := infrastructure.NewDB()
	userRepo := mysqlrepo.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	resolver := graph.NewResolver(userService)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: resolver,
	}))

	http.Handle("/query", srv)
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))

	server := &http.Server{
		Addr:    ":" + port,
		Handler: nil,
	}

	go func() {
		log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
}
