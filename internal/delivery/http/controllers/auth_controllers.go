package controllers

import (
	"net/http"

	"github.com/gene-qxsi/Flexive/internal/delivery/http/dto"
	"github.com/gene-qxsi/Flexive/internal/usecase"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	// AuthSrv     *services.AuthService
	AuthUseCase *usecase.AuthUseCase
}

func NewAuthController(authUseCase *usecase.AuthUseCase) *AuthController {
	return &AuthController{AuthUseCase: authUseCase}
}

func (ac *AuthController) SignIn(c *gin.Context) {
	// const op = "internal/delivery/http/controllers/auth_controllers.go/SignIn()"

	var req dto.SignInRequest
	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "не корректное тело запроса"})
		return
	}

	tokenResponse, err := ac.AuthUseCase.SignIn(c.Request.Context(), req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tokenResponse)
}

func (ac *AuthController) SignUp(c *gin.Context) {
	// const op = "internal/delivery/http/controllers/auth_controllers.go/SignUp()"

	var req dto.SignUpRequest
	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "не корректное тело запроса"})
		return
	}

	tokenResponse, err := ac.AuthUseCase.SignUp(c.Request.Context(), req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, tokenResponse)
}

func (ac *AuthController) SignOut(c *gin.Context) {
	// const op = "internal/delivery/http/controllers/auth_controllers.go/SignUp()"

	var req dto.RefreshToken
	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "не корректное тело запроса"})
		return
	}

	err := ac.AuthUseCase.SignOut(c.Request.Context(), req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (ac *AuthController) RefreshToken(c *gin.Context) {
	// const op = "internal/delivery/http/controllers/auth_controllers.go/RefreshToken()"

	var req dto.RefreshToken
	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "не корректное тело запроса"})
		return
	}

	tokenResponse, err := ac.AuthUseCase.RefreshToken(c.Request.Context(), req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tokenResponse)
}
