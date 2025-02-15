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

type ChannelHandler struct {
	service *services.ChannelService
}

func NewChannelHandler(service *services.ChannelService) *ChannelHandler {
	return &ChannelHandler{service: service}
}

func (h *ChannelHandler) CreateChannel(w http.ResponseWriter, r *http.Request) {
	const op = "internal/api/handlers/channel_handler.go/CreateChannel()"

	var channel *models.Channel
	err := json.NewDecoder(r.Body).Decode(&channel)
	if err != nil {
		err = fmt.Errorf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
		log.Println(err)
		utils.PushResponseJSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	_, err = h.service.CreateChannel(channel)
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

func (h *ChannelHandler) GetChannel(w http.ResponseWriter, r *http.Request) {
	const op = "internal/api/handlers/channel_handler.go/GetChannel()"

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		err = fmt.Errorf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
		log.Println(err)
		utils.PushResponseJSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	channel, err := h.service.GetChannel(id)
	if err != nil {
		err = fmt.Errorf("❌ ОБРАБОТЧИК-ОШИБКА-2: %s. ПУТЬ: %s", err.Error(), op)
		log.Println(err)
		utils.PushResponseJSON(w, http.StatusNotFound, err.Error(), nil)
		return
	}

	err = utils.PushResponseJSON(w, http.StatusOK, "ok", channel)
	if err != nil {
		err = fmt.Errorf("❌ ОБРАБОТЧИК-ОШИБКА-3: %s. ПУТЬ: %s", err.Error(), op)
		log.Println(err)
		return
	}

	log.Printf("✅ ОБРАБОТЧИК-УСПЕХ. ПУТЬ: %s\n", op)
}

func (h *ChannelHandler) GetChannels(w http.ResponseWriter, r *http.Request) {
	const op = "internal/api/handlers/channel_handler.go/GetChannels()"

	channels, err := h.service.GetChannels()
	if err != nil {
		err = fmt.Errorf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
		log.Println(err)
		utils.PushResponseJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	err = utils.PushResponseJSON(w, http.StatusOK, "ok", channels)
	if err != nil {
		err = fmt.Errorf("❌ ОБРАБОТЧИК-ОШИБКА-2: %s. ПУТЬ: %s", err.Error(), op)
		log.Println(err)
		return
	}

	log.Printf("✅ ОБРАБОТЧИК-УСПЕХ. ПУТЬ: %s\n", op)
}

func (h *ChannelHandler) UpdateChannel(w http.ResponseWriter, r *http.Request) {
	const op = "internal/api/handlers/channel_handler.go/UpdateChannel()"

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

	err = h.service.UpdateChannel(id, values)
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

func (h *ChannelHandler) DeleteChannel(w http.ResponseWriter, r *http.Request) {
	const op = "internal/api/handlers/channel_handler.go/DeleteChannel()"

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		err = fmt.Errorf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
		log.Println(err)
		utils.PushResponseJSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	err = h.service.DeleteChannel(id)
	if err != nil {
		err = fmt.Errorf("❌ ОБРАБОТЧИК-ОШИБКА-2: %s. ПУТЬ: %s", err.Error(), op)
		log.Println(err)
		utils.PushResponseJSON(w, http.StatusNotFound, err.Error(), nil)
		return
	}

	utils.PushResponseJSONNoContent(w)
	log.Printf("✅ ОБРАБОТЧИК-УСПЕХ. ПУТЬ: %s\n", op)
}
