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

type CommentHandler struct {
	service *services.CommentService
}

func NewCommentHandler(service *services.CommentService) *CommentHandler {
	return &CommentHandler{service: service}
}

func (h *CommentHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
	const op = "internal/api/handlers/comment_handler.go/CreateComment()"

	var comment *models.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		err = fmt.Errorf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
		log.Println(err)
		utils.PushResponseJSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	_, err = h.service.CreateComment(comment)
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

func (h *CommentHandler) GetComment(w http.ResponseWriter, r *http.Request) {
	const op = "internal/api/handlers/comment_handler.go/GetComment()"

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		err = fmt.Errorf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
		log.Println(err)
		utils.PushResponseJSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	comment, err := h.service.GetComment(id)
	if err != nil {
		err = fmt.Errorf("❌ ОБРАБОТЧИК-ОШИБКА-2: %s. ПУТЬ: %s", err.Error(), op)
		log.Println(err)
		utils.PushResponseJSON(w, http.StatusNotFound, err.Error(), nil)
		return
	}

	err = utils.PushResponseJSON(w, http.StatusOK, "ok", comment)
	if err != nil {
		err = fmt.Errorf("❌ ОБРАБОТЧИК-ОШИБКА-3: %s. ПУТЬ: %s", err.Error(), op)
		log.Println(err)
		return
	}

	log.Printf("✅ ОБРАБОТЧИК-УСПЕХ. ПУТЬ: %s\n", op)
}

func (h *CommentHandler) GetComments(w http.ResponseWriter, r *http.Request) {
	const op = "internal/api/handlers/comment_handler.go/GetComments()"

	comments, err := h.service.GetComments()
	if err != nil {
		err = fmt.Errorf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
		log.Println(err)
		utils.PushResponseJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	err = utils.PushResponseJSON(w, http.StatusOK, "ok", comments)
	if err != nil {
		err = fmt.Errorf("❌ ОБРАБОТЧИК-ОШИБКА-2: %s. ПУТЬ: %s", err.Error(), op)
		log.Println(err)
		return
	}

	log.Printf("✅ ОБРАБОТЧИК-УСПЕХ. ПУТЬ: %s\n", op)
}

func (h *CommentHandler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	const op = "internal/api/handlers/comment_handler.go/UpdateComment()"

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

	err = h.service.UpdateComment(id, values)
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

func (h *CommentHandler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	const op = "internal/api/handlers/comment_handler.go/DeleteComment()"

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		err = fmt.Errorf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
		log.Println(err)
		utils.PushResponseJSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	err = h.service.DeleteComment(id)
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
