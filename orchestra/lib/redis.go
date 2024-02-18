package lib

import (
	"fmt"
	"sync"

	"github.com/redis/go-redis/v9"
)

var (
	rdOnce sync.Once
	red    RedisWithMutex  // Для локального использования
	Rdb    *RedisWithMutex // Для глобального использования

	rdaddress = Getenv("REDIS_ADDRESS", "localhost:6379")
)

// RedisWithMutex Клиент Redis с мьютексом для защиты от гонки
type RedisWithMutex struct {
	Client *redis.Client
	Mu     *sync.RWMutex
}

// SetupRedis подключает Redis в первый раз
func SetupRedis() {
	// Инициализируем
	Rdb = NewRedisConnection()

	fmt.Println("SETUP: Redis database has been connected")
}

// NewRedisConnection Инициализация подключения к Redis базе данных.
// Возвращает ссылку на структуру RedisWithMutex
func NewRedisConnection() *RedisWithMutex {
	// Инициализируем только один раз
	rdOnce.Do(func() {
		var client = redis.NewClient(&redis.Options{
			Addr: rdaddress,
			DB:   0,
		})

		red.Client = client

		red.Mu = &sync.RWMutex{}
	})

	return &red
}
