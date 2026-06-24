package infrastructure

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/kojima1128/portfolio-backend/internal/common"
)

func NewDB() *gorm.DB {
	host := common.GetEnv("MYSQL_HOST", "localhost")
	port := common.GetEnv("MYSQL_PORT", "3306")
	user := common.GetEnv("MYSQL_USER", "user")
	password := common.GetEnv("MYSQL_PASSWORD", "password")
	database := common.GetEnv("MYSQL_DATABASE", "common_db")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, database,
	)

	var db *gorm.DB
	var err error
	for i := 0; i < 10; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("waiting for database... (%d/10): %v", i+1, err)
		time.Sleep(3 * time.Second)
	}
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	return db
}
