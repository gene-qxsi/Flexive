package storage

import (
	"database/sql"
	"fmt"

	"github.com/gene-qxsi/Flexive/configs"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	gPostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDB(gConf *gorm.Config, conf *configs.Config) (*gorm.DB, error) {
	const op = "internal/storage/db.go/NewStorage()"

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		conf.DbUser, conf.DbPassword, conf.DbName, conf.DbHost, conf.DbPort, conf.DbSslmode)

	db, err := gorm.Open(gPostgres.Open(dsn), gConf)
	if err != nil {
		return nil, fmt.Errorf("❌ БД-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}
	fmt.Println("УСПЕШНО ❌❌❌❌❌❌❌❌❌❌❌❌❌❌❌❌")
	err = runMigrations(dsn, conf)
	// db.AutoMigrate(models.User{}, models.Post{}, models.Channel{},
	// 	models.Comment{}, models.Subscription{})

	return db, err
}

func runMigrations(dsn string, conf *configs.Config) error {
	db, err := sql.Open(conf.Driver, dsn)
	if err != nil {
		return fmt.Errorf("ОШИБКА установки соединения бд: %s", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{
		DatabaseName: conf.DbName,
	})
	if err != nil {
		return fmt.Errorf("ОШИБКА создания инстанса: %s", err)
	}

	migrate, err := migrate.NewWithDatabaseInstance(conf.MigrationsPath, conf.DbName, driver)
	if err != nil {
		return fmt.Errorf("ОШИБКА создания объекта миграций: %s", err.Error())
	}

	if err = migrate.Up(); err != nil {
		return fmt.Errorf("ОШИБКА выполнения миграций: %s", err.Error())
	}

	return nil
}
