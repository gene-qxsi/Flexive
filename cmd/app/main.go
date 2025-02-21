package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gene-qxsi/Flexive/configs"
	"github.com/gene-qxsi/Flexive/internal/router"
	"github.com/gene-qxsi/Flexive/internal/services"
)

func main() {
	configs.Load()
	if err := services.Init(); err != nil {
		log.Fatal(err)
	}

	log.Println("🚀 Сервер запущен на", os.Getenv("GO_PORT"))
	http.ListenAndServe(os.Getenv("GO_PORT"), router.InitRouter())
}
