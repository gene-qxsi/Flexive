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

type PostHandler struct {
	service *services.PostService
}

func NewPostHandler(service *services.PostService) *PostHandler {
	return &PostHandler{service: service}
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	const op = "internal/api/handlers/post_handler.go/CreatePost()"

	var post *models.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		err = fmt.Errorf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
		log.Println(err)
		utils.PushResponseJSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	postDTO, err := h.service.CreatePost(post)
	if err != nil {
		err = fmt.Errorf("❌ ОБРАБОТЧИК-ОШИБКА-2: %s. ПУТЬ: %s", err.Error(), op)
		log.Println(err)
		utils.PushResponseJSON(w, http.StatusConflict, err.Error(), postDTO)
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

func (h *PostHandler) GetPost(w http.ResponseWriter, r *http.Request) {
	const op = "internal/api/handlers/post_handler.go/GetPost()"

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		err = fmt.Errorf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
		log.Println(err)
		utils.PushResponseJSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	post, err := h.service.GetPost(id)
	if err != nil {
		err = fmt.Errorf("❌ ОБРАБОТЧИК-ОШИБКА-2: %s. ПУТЬ: %s", err.Error(), op)
		log.Println(err)
		utils.PushResponseJSON(w, http.StatusNotFound, err.Error(), nil)
		return
	}

	err = utils.PushResponseJSON(w, http.StatusOK, "ok", post)
	if err != nil {
		err = fmt.Errorf("❌ ОБРАБОТЧИК-ОШИБКА-3: %s. ПУТЬ: %s", err.Error(), op)
		log.Println(err)
		return
	}

	log.Printf("✅ ОБРАБОТЧИК-УСПЕХ. ПУТЬ: %s\n", op)
}

func (h *PostHandler) GetPosts(w http.ResponseWriter, r *http.Request) {
	const op = "internal/api/handlers/post_handler.go/GetPosts()"

	posts, err := h.service.GetPosts()
	if err != nil {
		err = fmt.Errorf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
		log.Println(err)
		utils.PushResponseJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	err = utils.PushResponseJSON(w, http.StatusOK, "ok", posts)
	if err != nil {
		err = fmt.Errorf("❌ ОБРАБОТЧИК-ОШИБКА-2: %s. ПУТЬ: %s", err.Error(), op)
		log.Println(err)
		return
	}

	log.Printf("✅ ОБРАБОТЧИК-УСПЕХ. ПУТЬ: %s\n", op)
}

func (h *PostHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	const op = "internal/api/handlers/post_handler.go/UpdatePost()"

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

	postDTO, err := h.service.UpdatePost(id, values)
	if err != nil {
		err = fmt.Errorf("❌ ОБРАБОТЧИК-ОШИБКА-3: %s. ПУТЬ: %s", err.Error(), op)
		log.Println(err)
		utils.PushResponseJSON(w, http.StatusNotFound, err.Error(), nil)
		return
	}

	err = utils.PushResponseJSON(w, http.StatusOK, "ok", postDTO)
	if err != nil {
		err = fmt.Errorf("❌ ОБРАБОТЧИК-ОШИБКА-4: %s. ПУТЬ: %s", err.Error(), op)
		log.Println(err)
		return
	}

	log.Printf("✅ ОБРАБОТЧИК-УСПЕХ. ПУТЬ: %s\n", op)
}

func (h *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	const op = "internal/api/handlers/post_handler.go/DeletePost()"

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		err = fmt.Errorf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op)
		log.Println(err)
		utils.PushResponseJSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	err = h.service.DeletePost(id)
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
