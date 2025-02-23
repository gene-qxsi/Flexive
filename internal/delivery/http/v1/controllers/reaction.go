package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gene-qxsi/Flexive/internal/repository/models"
	"github.com/gene-qxsi/Flexive/internal/services"
	"github.com/gin-gonic/gin"
)

type ReactionHandler struct {
	service *services.ReactionService
}

func NewReactionHandler(service *services.ReactionService) *ReactionHandler {
	return &ReactionHandler{service: service}
}

func (h *ReactionHandler) CreateReaction(c *gin.Context) {
	const op = "internal/api/handlers/reaction_handler.go/CreateReaction()"

	var reaction *models.Reaction
	err := c.BindJSON(&reaction)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	reactionDTO, err := h.service.CreateReaction(reaction)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-2: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	c.JSON(http.StatusCreated, reactionDTO)
}

func (h *ReactionHandler) GetReaction(c *gin.Context) {
	const op = "internal/delivery/http/controllers/reaction_controller.go/GetReaction()"

	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	postID, err := strconv.Atoi(c.Param("postID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-2: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	reactionDTO, err := h.service.GetReaction(userID, postID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-3: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	c.JSON(http.StatusOK, reactionDTO)
}

func (h *ReactionHandler) GetReactions(c *gin.Context) {
	const op = "internal/delivery/http/controllers/reaction_controller.go/GetReactions()"

	reactionsDTOs, err := h.service.GetReactions()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	c.JSON(http.StatusOK, reactionsDTOs)
}

func (h *ReactionHandler) UpdateReaction(c *gin.Context) {
	const op = "internal/delivery/http/controllers/reaction_controller.go/UpdateReaction()"

	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	postID, err := strconv.Atoi(c.Param("postID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-2: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	var values map[string]interface{}
	err = c.BindJSON(&values)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-3: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	reactionDTO, err := h.service.UpdateReaction(userID, postID, values)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-4: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	c.JSON(http.StatusOK, reactionDTO)
}

func (h *ReactionHandler) DeleteReaction(c *gin.Context) {
	const op = "internal/delivery/http/controllers/reaction_controller.go/DeleteReaction()"

	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	postID, err := strconv.Atoi(c.Param("postID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-2: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	err = h.service.DeleteReaction(userID, postID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-3: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
