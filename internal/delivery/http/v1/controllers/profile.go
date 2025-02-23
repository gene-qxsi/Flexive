package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gene-qxsi/Flexive/internal/domain"
	"github.com/gene-qxsi/Flexive/internal/usecase"
	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	profileUsecase *usecase.ProfileUsecase
}

func NewProfileController(profileUsecase *usecase.ProfileUsecase) *ProfileController {
	return &ProfileController{profileUsecase: profileUsecase}
}

func (h *ProfileController) GetMyProfile(c *gin.Context) {
	const op = "internal/delivery/http/controllers/profile_controller.go/GetMyProfile()"

	claims, exists := c.Get("claims")
	fmt.Println(claims)
	if !exists {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", "токен не найден", op))
		return
	}

	userClaims, ok := claims.(*domain.AuthClaims)
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", "ошибка конвертации", op))
		return
	}

	resp, err := h.profileUsecase.GetProfileByUserID(c.Request.Context(), userClaims.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-2: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *ProfileController) GetProfile(c *gin.Context) {
	const op = "internal/delivery/http/controllers/profile_controller.go/GetProfile()"

	id, err := strconv.Atoi(c.Param("userID"))
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

func (h *ProfileController) GetProfiles(c *gin.Context) {
	const op = "internal/delivery/http/controllers/profile_controller.go/GetProfile()"

	resp, err := h.profileUsecase.GetProfiles(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *ProfileController) UpdateProfile(c *gin.Context) {
	const op = "internal/delivery/http/v1/controllers/profile_controller.go/UpdateProfile()"

	claims, exists := c.Get("claims")
	fmt.Println(claims)
	if !exists {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-1: %s. ПУТЬ: %s", "токен не найден", op))
		return
	}

	userClaims, ok := claims.(*domain.AuthClaims)
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-2: %s. ПУТЬ: %s", "ошибка конвертации", op))
		return
	}

	var values map[string]interface{}
	if err := c.BindJSON(&values); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-3: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	resp, err := h.profileUsecase.UpdateProfile(c.Request.Context(), userClaims.ID, values)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, fmt.Sprintf("❌ ОБРАБОТЧИК-ОШИБКА-4: %s. ПУТЬ: %s", err.Error(), op))
		return
	}

	c.JSON(http.StatusOK, resp)
}
