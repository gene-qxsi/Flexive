package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Load() {
	const op = "configs/config.go/Load()"

	err := godotenv.Load("configs/config.env")
	if err != nil {
		log.Printf("❌ КОНФИГ-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
		os.Exit(1)
	}
}
