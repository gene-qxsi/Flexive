package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/gene-qxsi/Flexive/configs"
	"github.com/redis/go-redis/v9"
)

type AuthRepository struct {
	client     *redis.Client
	refreshTTL time.Duration
}

func NewAuthRepository(client *redis.Client, cfg *configs.Config) *AuthRepository {
	return &AuthRepository{
		client:     client,
		refreshTTL: cfg.RedisRefreshTokenTTL,
	}
}

func (r *AuthRepository) SaveRefreshToken(ctx context.Context, token string, userID int) error {
	const op = "internal/repository/auth_repository.go/SaveRefreshToken()"

	err := r.client.Set(ctx, token, userID, r.refreshTTL).Err()
	if err != nil {
		return fmt.Errorf("ошибка сохранения токена. ПУТЬ: %s", op)
	}
	return nil
}

func (r *AuthRepository) GetUserIDByRefreshToken(ctx context.Context, token string) (int, error) {
	const op = "internal/repository/auth_repository.go/GetRefreshToken()"

	userID, err := r.client.Get(ctx, token).Int()
	if err == redis.Nil {
		return 0, redis.Nil
	}

	if err != nil {
		return 0, fmt.Errorf("%s: ошибка извлечения refresh токена: %w", op, err)
	}

	return userID, nil
}

func (r *AuthRepository) DeleteByRefreshToken(ctx context.Context, token string) error {
	const op = "internal/repository/auth_repository.go/DeleteRefreshToken"

	err := r.client.Del(ctx, token).Err()
	if err != nil {
		return fmt.Errorf("%s: ошибка удаления refresh-токена: %w", op, err)
	}
	return err
}
