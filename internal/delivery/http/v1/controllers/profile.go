package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gene-qxsi/Flexive/internal/usecase"
	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	profileUsecase *usecase.ProfileUsecase
}

func NewProfileController(profileUsecase *usecase.ProfileUsecase) *ProfileController {
	return &ProfileController{profileUsecase: profileUsecase}
}

func (h *ProfileController) GetProfile(c *gin.Context) {
	const op = "internal/delivery/http/controllers/profile_controller.go/GetProfile()"

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	resp, err := h.profileUsecase.GetProfileByUserID(c.Request.Context(), id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-2: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *ProfileController) UpdateProfile(c *gin.Context) {
	const op = "internal/delivery/http/controllers/profile_controller.go/UpdateProfile()"

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	var values map[string]string
	if err := c.BindJSON(&values); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-2: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	resp, err := h.profileUsecase.UpdateProfile(c.Request.Context(), id, values)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-3: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	c.JSON(http.StatusOK, resp)
}
