// Package common provides shared utilities used across the application.
package common

import (
	"fmt"
	"os"
)

// GetEnv returns the value of the environment variable named by key.
// If the variable is empty, it returns the provided default value.
func GetEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// DBConnectionString builds a MySQL DSN from environment variables.
func DBConnectionString() string {
	host := GetEnv("DB_HOST", "localhost")
	port := GetEnv("DB_PORT", "3306")
	user := GetEnv("DB_USER", "user")
	password := GetEnv("DB_PASSWORD", "password")
	name := GetEnv("DB_NAME", "common_db")
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4",
		user, password, host, port, name)
}
