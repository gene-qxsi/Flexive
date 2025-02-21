package cache

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	postsKey = "posts:ranking"
	ctx      = context.Background()
	ttl      = time.Hour * 1
)

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient() (*RedisClient, error) {
	const op = "internal/cache/redis_db.go/NewRedisClient()"

	// addr := os.Getenv("REDIS_ADDR")
	// password := os.Getenv("REDIS_PASSWORD")
	// db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

	// client := &RedisClient{
	// 	Client: redis.NewClient(&redis.Options{
	// 		Addr:     addr,
	// 		Password: password,
	// 		DB:       db,
	// 	}),
	// }

	client := &RedisClient{
		Client: redis.NewClient(&redis.Options{}),
	}

	_, err := client.Client.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("❌ КЕШ-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	log.Printf("✅ КЕШ-УСПЕХ. ПУТЬ: %s\n", op)
	return client, nil
}

func NewRedisClientCustom(addr, password string, db int) (*RedisClient, error) {
	const op = "internal/cache/redis_db.go/NewRedisClientCustom()"

	client := &RedisClient{
		Client: redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
			DB:       db,
		}),
	}

	_, err := client.Client.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("❌ КЕШ-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	log.Printf("✅ КЕШ-УСПЕХ. ПУТЬ: %s\n", op)
	return client, nil
}

func (r *RedisClient) IncrPost(member string) error {
	const op = "internal/cache/redis_db.go/CachePost()"

	_, err := r.Client.ZIncrBy(ctx, postsKey, 1, member).Result()
	if err != nil {
		return fmt.Errorf("❌ КЕШ-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	_, err = r.Client.Expire(ctx, postsKey, ttl).Result()
	if err != nil {
		return fmt.Errorf("❌ КЕШ-ОШИБКА-2: %s. ПУТЬ: %s", err.Error(), op)
	}

	log.Printf("✅ КЕШ-УСПЕХ. ПУТЬ: %s\n", op)
	return nil
}

func (r *RedisClient) GetTopPostsIDs(limit int) ([]int, error) {
	const op = "internal/cache/redis_db.go/GetTopN()"

	idS, err := r.Client.ZRevRange(ctx, postsKey, 0, int64(limit-1)).Result()
	if err != nil {
		return nil, fmt.Errorf("❌ КЕШ-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	var idIs []int
	for _, id := range idS {
		idI, err := strconv.Atoi(id)
		if err != nil {
			return nil, fmt.Errorf("❌ КЕШ-ОШИБКА-2: %s. ПУТЬ: %s", err, op)
		}
		idIs = append(idIs, idI)
	}

	log.Printf("✅ КЕШ-УСПЕХ. ПУТЬ: %s\n", op)
	return idIs, nil
}
