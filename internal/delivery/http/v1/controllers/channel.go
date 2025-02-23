package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gene-qxsi/Flexive/internal/repository/models"
	"github.com/gene-qxsi/Flexive/internal/services"
	"github.com/gin-gonic/gin"
)

type ChannelController struct {
	service *services.ChannelService
}

func NewChannelController(service *services.ChannelService) *ChannelController {
	return &ChannelController{service: service}
}

func (h *ChannelController) CreateChannel(c *gin.Context) {
	const op = "internal/delivery/http/controllers/channel_controller.go/CreateChannel()"

	var channel *models.Channel
	err := c.BindJSON(&channel)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	commentDTO, err := h.service.CreateChannel(channel)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-2: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	c.JSON(http.StatusCreated, commentDTO)
}

func (h *ChannelController) GetChannel(c *gin.Context) {
	const op = "internal/delivery/http/controllers/channel_controller.go/GetChannel()"

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	channelDTO, err := h.service.GetChannel(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-2: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	c.JSON(http.StatusOK, channelDTO)
}

func (h *ChannelController) GetChannels(c *gin.Context) {
	const op = "internal/delivery/http/controllers/channel_controller.go/GetChannels()"

	channelsDTOs, err := h.service.GetChannels()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	c.JSON(http.StatusOK, channelsDTOs)
}

func (h *ChannelController) UpdateChannel(c *gin.Context) {
	const op = "internal/delivery/http/controllers/channel_controller.go/UpdateChannel()"

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

	channelDTO, err := h.service.UpdateChannel(id, values)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-3: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	c.JSON(http.StatusOK, channelDTO)
}

func (h *ChannelController) DeleteChannel(c *gin.Context) {
	const op = "internal/delivery/http/controllers/channel_controller.go/DeleteChannel()"

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	err = h.service.DeleteChannel(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-2: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
