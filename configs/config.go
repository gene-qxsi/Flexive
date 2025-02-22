package configs

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	RedisRefreshTokenTTL time.Duration
	RedisAccessTokenTTL  time.Duration
	PostsRankingTTL      time.Duration
	Salt                 string
	GoServerAddr         string
	JWTSecretKey         string
	DbUser               string
	DbPassword           string
	DbPort               string
	DbHost               string
	DbName               string
	DbSslmode            string
}

func Load() *Config {
	const op = "configs/config.go/Load()"

	err := godotenv.Load("configs/config.env")
	if err != nil {
		log.Printf("❌ КОНФИГ-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
		os.Exit(1)
	}

	rfhTTL, err := strconv.Atoi(os.Getenv("REFRESH_TOKEN_TTL"))
	if err != nil {
		rfhTTL = 44640
	}

	acsTTL, err := strconv.Atoi(os.Getenv("ACCESS_TOKEN_TTL"))
	if err != nil {
		acsTTL = 15
	}

	postsRankingTTL, err := strconv.Atoi(os.Getenv("POSTS_RANKING_TTL"))
	if err != nil {
		postsRankingTTL = 60
	}

	goServerAddr := os.Getenv("GO_ADDR")
	if err != nil {
		goServerAddr = ":8080"
	}

	salt := os.Getenv("SALT")
	if err != nil {
		salt = "flexive2202"
	}

	JWTSecretKey := os.Getenv("JWT_SECRET_KEY")
	if err != nil {
		JWTSecretKey = "flexive1702"
	}

	dbUser := os.Getenv("DB_USER")
	if err != nil {
		dbUser = "postgres"
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if err != nil {
		dbPassword = "admin"
	}

	dbName := os.Getenv("DB_NAME")
	if err != nil {
		dbName = "moji"
	}

	dbHost := os.Getenv("DB_HOST")
	if err != nil {
		dbHost = "localhost"
	}

	dbPort := os.Getenv("DB_PORT")
	if err != nil {
		dbPort = "5432"
	}

	dbSslmode := os.Getenv("DB_SSLMODE")
	if err != nil {
		dbSslmode = "disable"
	}

	return &Config{
		RedisRefreshTokenTTL: time.Duration(rfhTTL) * time.Minute,
		RedisAccessTokenTTL:  time.Duration(acsTTL) * time.Minute,
		PostsRankingTTL:      time.Duration(postsRankingTTL) * time.Minute,
		Salt:                 salt,
		GoServerAddr:         goServerAddr,
		JWTSecretKey:         JWTSecretKey,
		DbUser:               dbUser,
		DbPassword:           dbPassword,
		DbPort:               dbPort,
		DbHost:               dbHost,
		DbName:               dbName,
		DbSslmode:            dbSslmode,
	}
}
