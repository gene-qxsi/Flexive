package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gene-qxsi/Flexive/internal/domain"
	"github.com/gin-gonic/gin"
)

type ChatUsecase interface {
	GetChats(ctx context.Context) ([]domain.Chat, error)
	CreateChat(ctx context.Context, chat domain.Chat) (*domain.Chat, error)
}

type ChatController struct {
	ChatUsecase ChatUsecase
}

func NewChatController(chatUsecase ChatUsecase) *ChatController {
	return &ChatController{ChatUsecase: chatUsecase}
}

func (h *ChatController) GetChats(c *gin.Context) {
	// const op = "interna/controllers/http/v1/chat.go/GetChats()"

	chatsDTOs, err := h.ChatUsecase.GetChats(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "непонятная ошибка, вот: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, chatsDTOs)
}

func (h *ChatController) CreateChat(c *gin.Context) {
	// const op = "interna/controllers/http/v1/chat.go/GetChats()"

	claims, exists := c.Get("claims")
	fmt.Println(claims)
	if !exists {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "пользоатель не прошел аутентификацию"})
		return
	}

	userClaims, ok := claims.(*domain.AuthClaims)
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ошибка преобразования claims -> *domain.AuthClaims"})
		return
	}

	var chat domain.Chat
	if err := c.BindJSON(&chat); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "не получилось распарсить JSON в chat: " + err.Error()})
		return
	}

	chat.UserID = userClaims.ID

	chatDTO, err := h.ChatUsecase.CreateChat(c.Request.Context(), chat)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "еще непонятная ошибка: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, chatDTO)
}
