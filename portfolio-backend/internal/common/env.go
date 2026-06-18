package common

import "os"

// GetEnv は環境変数を取得します。値が未設定の場合は defaultValue を返します。
func GetEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
