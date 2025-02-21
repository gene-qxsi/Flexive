package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gene-qxsi/Flexive/internal/repository/models"
	"github.com/gene-qxsi/Flexive/internal/services"
	"github.com/gin-gonic/gin"
)

type SubscriptionController struct {
	service *services.SubscriptionService
}

func NewSubscriptionController(service *services.SubscriptionService) *SubscriptionController {
	return &SubscriptionController{service: service}
}

func (h *SubscriptionController) CreateSubscription(c *gin.Context) {
	const op = "internal/delivery/http/controllers/subscription_controller.go/CreateSubscription()"

	var subscription *models.Subscription
	err := c.BindJSON(&subscription)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	subscriptionDTO, err := h.service.CreateSubscription(subscription)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-2: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	c.JSON(http.StatusCreated, subscriptionDTO)
}

func (h *SubscriptionController) GetSubscription(c *gin.Context) {
	const op = "internal/delivery/http/controllers/subscription_controller.go/GetSubscription()"

	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	channelID, err := strconv.Atoi(c.Param("channelID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-2: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	subscriptionDTO, err := h.service.GetSubscription(userID, channelID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-3: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	c.JSON(http.StatusOK, subscriptionDTO)
}

func (h *SubscriptionController) GetSubscriptions(c *gin.Context) {
	const op = "internal/delivery/http/controllers/subscription_controller.go/GetSubscriptions()"

	subscriptionsDTOs, err := h.service.GetSubscriptions()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	c.JSON(http.StatusOK, subscriptionsDTOs)
}

func (h *SubscriptionController) UpdateSubscription(c *gin.Context) {
	const op = "internal/delivery/http/controllers/subscription_controller.go/UpdateSubscription()"

	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	channelID, err := strconv.Atoi(c.Param("channelID"))
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

	subscriptionDTO, err := h.service.UpdateSubscription(userID, channelID, values)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-4: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	c.JSON(http.StatusOK, subscriptionDTO)
}

func (h *SubscriptionController) DeleteSubscription(c *gin.Context) {
	const op = "internal/delivery/http/controllers/subscription_controller.go/DeleteSubscription()"

	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	channelID, err := strconv.Atoi(c.Param("channelID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-2: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	err = h.service.DeleteSubscription(userID, channelID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-3: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
