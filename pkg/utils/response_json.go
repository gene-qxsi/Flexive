package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gene-qxsi/Flexive/internal/models/http_models"
)

func PushResponseJSON(w http.ResponseWriter, status int, message string, data interface{}) error {
	const op = "utils/response_json.go/PushResponseJSON()"

	response := http_models.HTTPResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return fmt.Errorf("❌ РЕПОЗИТОРИЙ-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
	}
	return nil
}

func PushResponseJSONNoContent(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
