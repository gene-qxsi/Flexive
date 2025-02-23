package storage

import (
	"fmt"

	"github.com/gene-qxsi/Flexive/configs"
	"github.com/gene-qxsi/Flexive/internal/repository/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDB(gConf *gorm.Config, conf *configs.Config) (*gorm.DB, error) {
	const op = "internal/storage/db.go/NewStorage()"

	sdn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		conf.DbUser, conf.DbPassword, conf.DbName, conf.DbHost, conf.DbPort, conf.DbSslmode)

	db, err := gorm.Open(postgres.Open(sdn), gConf)
	if err != nil {
		return nil, fmt.Errorf("❌ БД-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}

	db.AutoMigrate(models.User{}, models.Post{}, models.Channel{},
		models.Comment{}, models.Subscription{})

	return db, nil
}
