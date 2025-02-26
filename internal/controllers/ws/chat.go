package controllers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/gene-qxsi/Flexive/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type ChatUsecase interface {
	LoadLastNMessagesByChatID(ctx context.Context, chatID, count int) ([]domain.Message, error)
	SaveMessage(ctx context.Context, chatID int, message domain.Message) error
}

type WSController struct {
	ChatUsecase     ChatUsecase
	chatConnections map[int][]*websocket.Conn
	mu              sync.Mutex
}

func NewWSController(chat ChatUsecase) *WSController {
	return &WSController{
		ChatUsecase:     chat,
		chatConnections: make(map[int][]*websocket.Conn),
	}
}

func (ws *WSController) ChatController(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Ошибка при апгрейде соединения:", err)
		return
	}
	defer conn.Close()

	chatId := c.DefaultQuery("id", "")
	if chatId == "" {
		fmt.Println("Ошибка: ID чата не указан")
		return
	}

	id, _ := strconv.Atoi(chatId)

	messages, err := ws.ChatUsecase.LoadLastNMessagesByChatID(c.Request.Context(), id, 20)
	if err != nil {
		fmt.Println("Ошибка загрузки сообщений:", err)
		return
	}

	if err := conn.WriteJSON(messages); err != nil {
		fmt.Println("Ошибка отправки истории чата:", err)
		return
	}

	ws.mu.Lock()
	ws.chatConnections[id] = append(ws.chatConnections[id], conn)
	ws.mu.Unlock()

	for {
		var message domain.Message
		err := conn.ReadJSON(&message)
		if err != nil {
			fmt.Println("Ошибка чтения:", err)
			break
		}

		go func() {
			if err := ws.ChatUsecase.SaveMessage(c.Request.Context(), id, message); err != nil {
				fmt.Println("Ошибка сохранения сообщения:", err)
			}
		}()

		ws.mu.Lock()
		for _, connection := range ws.chatConnections[id] {
			if err := connection.WriteJSON(message); err != nil {
				fmt.Println("Ошибка отправки сообщения:", err)
			}
		}
		ws.mu.Unlock()
	}
}
