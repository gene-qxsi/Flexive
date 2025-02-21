package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gene-qxsi/Flexive/internal/repository/models"
	"github.com/gene-qxsi/Flexive/internal/services"
	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	service *services.CommentService
}

func NewCommentHandler(service *services.CommentService) *CommentHandler {
	return &CommentHandler{service: service}
}

func (h *CommentHandler) CreateComment(c *gin.Context) {
	const op = "internal/delivery/http/controllers/comment_controller.go/CreateComment()"

	var comment *models.Comment
	err := c.BindJSON(&comment)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	commentDTO, err := h.service.CreateComment(comment)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-2: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	c.JSON(http.StatusCreated, commentDTO)
}

func (h *CommentHandler) GetComment(c *gin.Context) {
	const op = "internal/delivery/http/controllers/comment_controller.go/GetComment()"

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	commentDTO, err := h.service.GetComment(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-2: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	c.JSON(http.StatusOK, commentDTO)
}

func (h *CommentHandler) GetComments(c *gin.Context) {
	const op = "internal/delivery/http/controllers/comment_controller.go/GetComments()"

	commentsDTOs, err := h.service.GetComments()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	c.JSON(http.StatusOK, commentsDTOs)
}

func (h *CommentHandler) UpdateComment(c *gin.Context) {
	const op = "internal/delivery/http/controllers/comment_controller.go/UpdateComment()"

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	var values map[string]interface{}
	err = c.BindJSON(&values)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-2: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	commentDTO, err := h.service.UpdateComment(id, values)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-3: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	c.JSON(http.StatusOK, commentDTO)
}

func (h *CommentHandler) DeleteComment(c *gin.Context) {
	const op = "internal/delivery/http/controllers/comment_controller.go/DeleteComment()"

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	err = h.service.DeleteComment(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-2: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
