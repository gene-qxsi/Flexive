package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gene-qxsi/Flexive/internal/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserController(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	const op = "internal/delivery/http/controllers/user_controller.go/GetUser()"

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	userDTO, err := h.service.GetUser(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-2: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	c.JSON(http.StatusOK, userDTO)
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	const op = "internal/delivery/http/controllers/user_controller.go/GetUsers()"

	usersDTOs, err := h.service.GetUsers()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	c.JSON(http.StatusOK, usersDTOs)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	const op = "internal/delivery/http/controllers/user_controller.go/UpdateUser()"

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

	userDTO, err := h.service.UpdateUser(id, values)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-3: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	c.JSON(http.StatusOK, userDTO)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	const op = "internal/delivery/http/controllers/user_controller.go/DeleteUser()"

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	err = h.service.DeleteUser(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-2: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// func (h *UserHandler) CreateUser(c *gin.Context) {
// 	const op = "internal/delivery/http/controllers/user_controller.go/CreateUser()"

// 	var user *dto.SignUpRequest
// 	err := c.BindJSON(&user)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op))
// 		return
// 	}
// 	//TODO: нада какта эта исправить
// 	//  не нада передавать domain из controller слоя
// 	//  мб сделать прокси usecase
// 	userDTO, err := h.service.CreateUser(&domain.User{
// 		Username: user.Username,
// 		Email:    user.Email,
// 		Password: user.Password,
// 		Role:     user.Role,
// 		Birthday: user.Birthday,
// 	})
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusConflict, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-2: %s. ПУТЬ: %s", err.Error(), op))
// 		return
// 	}

// 	c.JSON(http.StatusCreated, userDTO)
// }
