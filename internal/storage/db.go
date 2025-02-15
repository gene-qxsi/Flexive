package storage

import (
	"fmt"
	"os"

	"github.com/gene-qxsi/Flexive/internal/cache"
	models "github.com/gene-qxsi/Flexive/internal/models/orm_models"
	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	Sdb *gorm.DB
	Rdb *cache.RedisClient
}

func NewStorage() (*Storage, error) {
	const op = "internal/storage/db.go/NewStorage()"

	err := godotenv.Load("config.env")
	if err != nil {
		return nil, fmt.Errorf("❌ БД-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	sdbConnStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"))

	db, err := gorm.Open(postgres.Open(sdbConnStr), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("❌ БД-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	db.AutoMigrate(models.User{}, models.Post{}, models.Channel{},
		models.Comment{}, models.Reaction{}, models.Subscription{})

	rdb, err := cache.NewRedisClient()
	if err != nil {
		return nil, fmt.Errorf("❌ БД-ОШИБКА-2: %s. ПУТЬ: %s", err.Error(), op)
	}

	return &Storage{Sdb: db, Rdb: rdb}, nil
}
