package main

import (
	"log"
	"net/http"

	"github.com/gene-qxsi/Flexive/configs"
	"github.com/gene-qxsi/Flexive/internal/router"
)

func main() {
	conf := configs.Load()

	mux := router.InitRouter(conf)

	log.Println("🚀 Сервер запущен на", conf.GoServerAddr)
	http.ListenAndServe(conf.GoServerAddr, mux)
}
