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

type SubscriptionHandler struct {
	service *services.SubscriptionService
}

func NewSubscriptionHandler(service *services.SubscriptionService) *SubscriptionHandler {
	return &SubscriptionHandler{service: service}
}

func (h *SubscriptionHandler) CreateSubscription(w http.ResponseWriter, r *http.Request) {
	const op = "internal/api/handlers/subscription_handler.go/CreateSubscription()"

	var subscription *models.Subscription
	err := json.NewDecoder(r.Body).Decode(&subscription)
	if err != nil {
		err = fmt.Errorf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
		log.Println(err)
		utils.PushResponseJSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	_, _, err = h.service.CreateSubscription(subscription)
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

func (h *SubscriptionHandler) GetSubscription(w http.ResponseWriter, r *http.Request) {
	const op = "internal/api/handlers/subscription_handler.go/GetSubscription()"

	userID, err := strconv.Atoi(chi.URLParam(r, "userID"))
	if err != nil {
		err = fmt.Errorf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
		log.Println(err)
		utils.PushResponseJSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	channelID, err := strconv.Atoi(chi.URLParam(r, "channelID"))
	if err != nil {
		err = fmt.Errorf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
		log.Println(err)
		utils.PushResponseJSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	reaction, err := h.service.GetSubscription(userID, channelID)
	if err != nil {
		err = fmt.Errorf("❌ ОБРАБОТЧИК-ОШИБКА-2: %s. ПУТЬ: %s", err.Error(), op)
		log.Println(err)
		utils.PushResponseJSON(w, http.StatusNotFound, err.Error(), nil)
		return
	}

	err = utils.PushResponseJSON(w, http.StatusOK, "ok", reaction)
	if err != nil {
		err = fmt.Errorf("❌ ОБРАБОТЧИК-ОШИБКА-3: %s. ПУТЬ: %s", err.Error(), op)
		log.Println(err)
		return
	}

	log.Printf("✅ ОБРАБОТЧИК-УСПЕХ. ПУТЬ: %s\n", op)
}

func (h *SubscriptionHandler) GetSubscriptions(w http.ResponseWriter, r *http.Request) {
	const op = "internal/api/handlers/subscription_handler.go/GetSubscription()"

	subscriptions, err := h.service.GetSubscriptions()
	if err != nil {
		err = fmt.Errorf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
		log.Println(err)
		utils.PushResponseJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	err = utils.PushResponseJSON(w, http.StatusOK, "ok", subscriptions)
	if err != nil {
		err = fmt.Errorf("❌ ОБРАБОТЧИК-ОШИБКА-2: %s. ПУТЬ: %s", err.Error(), op)
		log.Println(err)
		return
	}

	log.Printf("✅ ОБРАБОТЧИК-УСПЕХ. ПУТЬ: %s\n", op)
}

func (h *SubscriptionHandler) UpdateSubscription(w http.ResponseWriter, r *http.Request) {
	const op = "internal/api/handlers/subscription_handler.go/UpdateSubscription()"

	userID, err := strconv.Atoi(chi.URLParam(r, "userID"))
	if err != nil {
		err = fmt.Errorf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
		log.Println(err)
		utils.PushResponseJSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	channelID, err := strconv.Atoi(chi.URLParam(r, "channelID"))
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

	err = h.service.UpdateSubscription(userID, channelID, values)
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

func (h *SubscriptionHandler) DeleteSubscription(w http.ResponseWriter, r *http.Request) {
	const op = "internal/api/handlers/subscription_handler.go/DeleteSubscription()"

	userID, err := strconv.Atoi(chi.URLParam(r, "userID"))
	if err != nil {
		err = fmt.Errorf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
		log.Println(err)
		utils.PushResponseJSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	channelID, err := strconv.Atoi(chi.URLParam(r, "channelID"))
	if err != nil {
		err = fmt.Errorf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
		log.Println(err)
		utils.PushResponseJSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	err = h.service.DeleteSubscription(userID, channelID)
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
