package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kojima1128/portfolio-backend/internal/common"
	"github.com/kojima1128/portfolio-backend/internal/graph"
	"github.com/kojima1128/portfolio-backend/internal/service"
)

func main() {
	port := common.GetEnv("PORT", "8080")

	// TODO: DB接続を初期化し、リポジトリ・サービス層を構築する
	// db, err := mysql.Open(...)
	// userRepo := mysqlrepo.NewUserRepository(db)
	userSvc := service.NewUserService(nil) // nilはDB実装が完成するまでのプレースホルダー

	resolver := &graph.Resolver{
		UserService: userSvc,
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: resolver,
	}))

	http.Handle("/query", srv)
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
