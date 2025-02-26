package domain

import (
	"time"

	"github.com/gene-qxsi/Flexive/internal/repository/models"
)

type Chat struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	UserID    int       `json:"userID"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func ToDomainChat(orm *models.Chat) *Chat {
	return &Chat{
		ID:        orm.ID,
		Title:     orm.Title,
		UserID:    orm.UserID,
		CreatedAt: orm.CreatedAt,
		UpdatedAt: orm.UpdatedAt,
	}
}

func ToORMChat(chat *Chat) *models.Chat {
	return &models.Chat{
		ID:        chat.ID,
		Title:     chat.Title,
		UserID:    chat.UserID,
		CreatedAt: chat.CreatedAt,
		UpdatedAt: chat.UpdatedAt,
	}
}
