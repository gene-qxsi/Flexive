package storage

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

func OpenRedis(opt *redis.Options) (*redis.Client, error) {
	const op = "internal/storage/db.go/OpenRedis()"
	// пока так
	client := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("❌ КЕШ-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	return client, nil
}
