package main

import (
	"net/http"

	"github.com/gene-qxsi/Flexive/internal/api/router"
)

func main() {
	http.ListenAndServe(":8080", router.InitRouter())
}
