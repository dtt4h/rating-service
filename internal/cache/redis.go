package cache

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

// Создает и возвращает Redis клиенты
func NewRedisClient() *redis.Client {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return rdb
}

// Возвращает существующий клиент
func GetRedisClient() *redis.Client {
	if rdb == nil {
		rdb = NewRedisClient()
	}
	return rdb
}

// Проверяет подключение к Redis
func Ping() error {
	client := GetRedisClient()
	return client.Ping(ctx).Err()
}
