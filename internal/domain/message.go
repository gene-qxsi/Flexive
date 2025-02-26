package domain

import (
	"time"

	"github.com/gene-qxsi/Flexive/internal/repository/models"
)

type Message struct {
	ID        int       `json:"id"`
	ChatID    int       `json:"chat_id"`
	UserID    int       `json:"user_id"`
	Username  string    `json:"username"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToDomainMessage(orm *models.Message) *Message {
	return &Message{
		ID:        orm.ID,
		UserID:    orm.UserID,
		ChatID:    orm.ChatID,
		Content:   orm.Content,
		CreatedAt: orm.CreatedAt,
		UpdatedAt: orm.UpdatedAt,
	}
}

func ToORMMessage(message *Message) *models.Message {
	return &models.Message{
		ID:        message.ID,
		UserID:    message.UserID,
		ChatID:    message.ChatID,
		Content:   message.Content,
		CreatedAt: message.CreatedAt,
		UpdatedAt: message.UpdatedAt,
	}
}
