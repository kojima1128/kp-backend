package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/kojima1128/portfolio-backend/internal/common"
)

func main() {
	log.Printf("Starting worker (interval: %s)", common.GetEnv("WORKER_INTERVAL", "60s"))

	// TODO: DB接続を初期化し、リポジトリ・サービス層を構築する
	// TODO: ワーカーのジョブを実装する

	// グレースフルシャットダウンのためのシグナル待機
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Worker shutting down...")
}
