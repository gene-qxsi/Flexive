package storage

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func OpenRedis(opt *redis.Options) (*redis.Client, error) {
	const op = "internal/storage/db.go/OpenRedis()"

	client := redis.NewClient(opt)

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("❌ КЕШ-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	return client, nil
}
