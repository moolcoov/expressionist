package lib

import "os"

// Getenv возвращает значение из окружения или defaultValue
func Getenv(key string, defaultValue string) string {
	value, found := os.LookupEnv(key)
	if !found {
		return defaultValue
	}
	return value
}
