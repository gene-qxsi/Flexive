package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gene-qxsi/Flexive/internal/api/services"
	models "github.com/gene-qxsi/Flexive/internal/models/orm_models"
	"github.com/gene-qxsi/Flexive/pkg/utils"
	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	const op = "internal/api/handlers/user_handler.go/CreateUser()"

	var user *models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		err = fmt.Errorf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
		log.Println(err)
		utils.PushResponseJSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	_, err = h.service.CreateUser(user)
	if err != nil {
		err = fmt.Errorf("❌ ОБРАБОТЧИК-ОШИБКА-2: %s. ПУТЬ: %s", err.Error(), op)
		log.Println(err)
		utils.PushResponseJSON(w, http.StatusConflict, err.Error(), nil)
		return
	}

	err = utils.PushResponseJSON(w, http.StatusCreated, "ok", nil)
	if err != nil {
		err = fmt.Errorf("❌ ОБРАБОТЧИК-ОШИБКА-3: %s. ПУТЬ: %s", err.Error(), op)
		log.Println(err)
		return
	}
	log.Printf("✅ ОБРАБОТЧИК-УСПЕХ. ПУТЬ: %s\n", op)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	const op = "internal/api/handlers/user_handler.go/GetUser()"

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		err = fmt.Errorf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
		log.Println(err)
		utils.PushResponseJSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	user, err := h.service.GetUser(id)
	if err != nil {
		err = fmt.Errorf("❌ ОБРАБОТЧИК-ОШИБКА-2: %s. ПУТЬ: %s", err.Error(), op)
		log.Println(err)
		utils.PushResponseJSON(w, http.StatusNotFound, err.Error(), nil)
		return
	}

	err = utils.PushResponseJSON(w, http.StatusOK, "ok", user)
	if err != nil {
		err = fmt.Errorf("❌ ОБРАБОТЧИК-ОШИБКА-3: %s. ПУТЬ: %s", err.Error(), op)
		log.Println(err)
		return
	}

	log.Printf("✅ ОБРАБОТЧИК-УСПЕХ. ПУТЬ: %s\n", op)
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	const op = "internal/api/handlers/user_handler.go/GetUsers()"

	users, err := h.service.GetUsers()
	if err != nil {
		err = fmt.Errorf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
		log.Println(err)
		utils.PushResponseJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	err = utils.PushResponseJSON(w, http.StatusOK, "ok", users)
	if err != nil {
		err = fmt.Errorf("❌ ОБРАБОТЧИК-ОШИБКА-2: %s. ПУТЬ: %s", err.Error(), op)
		log.Println(err)
		return
	}

	log.Printf("✅ ОБРАБОТЧИК-УСПЕХ. ПУТЬ: %s\n", op)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	const op = "internal/api/handlers/user_handler.go/UpdateUser()"

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		err = fmt.Errorf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
		log.Println(err)
		utils.PushResponseJSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	var values map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&values)
	if err != nil {
		err = fmt.Errorf("❌ ОБРАБОТЧИК-ОШИБКА-2: %s. ПУТЬ: %s", err.Error(), op)
		log.Println(err)
		utils.PushResponseJSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	err = h.service.UpdateUser(id, values)
	if err != nil {
		err = fmt.Errorf("❌ ОБРАБОТЧИК-ОШИБКА-3: %s. ПУТЬ: %s", err.Error(), op)
		log.Println(err)
		utils.PushResponseJSON(w, http.StatusNotFound, err.Error(), nil)
		return
	}

	err = utils.PushResponseJSON(w, http.StatusOK, "ok", values)
	if err != nil {
		err = fmt.Errorf("❌ ОБРАБОТЧИК-ОШИБКА-4: %s. ПУТЬ: %s", err.Error(), op)
		log.Println(err)
		return
	}

	log.Printf("✅ ОБРАБОТЧИК-УСПЕХ. ПУТЬ: %s\n", op)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	const op = "internal/api/handlers/user_handler.go/DeleteUser()"

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		err = fmt.Errorf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
		log.Println(err)
		utils.PushResponseJSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	err = h.service.DeleteUser(id)
	if err != nil {
		err = fmt.Errorf("❌ ОБРАБОТЧИК-ОШИБКА-2: %s. ПУТЬ: %s", err.Error(), op)
		log.Println(err)
		utils.PushResponseJSON(w, http.StatusNotFound, err.Error(), nil)
		return
	}

	err = utils.PushResponseJSON(w, http.StatusNoContent, "ok", nil)
	if err != nil {
		err = fmt.Errorf("❌ ОБРАБОТЧИК-ОШИБКА-3: %s. ПУТЬ: %s", err.Error(), op)
		log.Println(err)
		return
	}

	log.Printf("✅ ОБРАБОТЧИК-УСПЕХ. ПУТЬ: %s\n", op)
}
