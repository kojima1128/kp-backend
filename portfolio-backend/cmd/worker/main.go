package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
	"time"

	"github.com/kojima1128/portfolio-backend/internal/common"
	"github.com/kojima1128/portfolio-backend/internal/repository"
	"github.com/kojima1128/portfolio-backend/internal/service"
)

func main() {
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	_ = userService

	log.Println("worker started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = ctx
	log.Println("worker stopped")
	_ = common.GetEnv("WORKER_MODE", "default")
}
