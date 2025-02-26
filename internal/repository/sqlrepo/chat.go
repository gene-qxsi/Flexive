package repository

import (
	"context"
	"fmt"

	"github.com/gene-qxsi/Flexive/internal/repository/models"
	"gorm.io/gorm"
)

type ChatRepo struct {
	db *gorm.DB
}

func NewChatRepo(db *gorm.DB) *ChatRepo {
	return &ChatRepo{db: db}
}

func (r *ChatRepo) GetChats(ctx context.Context) ([]models.Chat, error) {
	const op = "internal/repositry/sqlrepo/chat.go/GetChats()"

	var chats []models.Chat
	if err := r.db.WithContext(ctx).Preload("User").Debug().Find(&chats).Error; err != nil {
		return nil, fmt.Errorf("ОШИБКА: %s. ПУТЬ: %s", err.Error(), op)
	}

	return chats, nil
}

func (r *ChatRepo) CreateChat(ctx context.Context, chat *models.Chat) (*models.Chat, error) {
	const op = "internal/repositry/sqlrepo/chat.go/CreateChat()"

	if err := r.db.Model(&chat).WithContext(ctx).Debug().Create(&chat).Error; err != nil {
		return nil, fmt.Errorf("ОШИБКА: %s. ПУТЬ: %s", err.Error(), op)
	}

	return chat, nil
}

func (r *ChatRepo) GetMessagesByChatID(ctx context.Context, chatID, limit int) ([]models.Message, error) {
	const op = "internal/repositry/sqlrepo/chat.go/GetMessagesByChatID()"

	var messages []models.Message
	if err := r.db.WithContext(ctx).Debug().Limit(limit).Find(&messages).Error; err != nil {
		return nil, fmt.Errorf("ОШИБКА: %s. ПУТЬ: %s", err.Error(), op)
	}

	return messages, nil
}

func (r *ChatRepo) CreateMessage(ctx context.Context, chatID int, message *models.Message) error {
	const op = "internal/repositry/sqlrepo/chat.go/CreateMessage()"

	if err := r.db.Model(&message).WithContext(ctx).Debug().Create(&message).Error; err != nil {
		return fmt.Errorf("ОШИБКА: %s. ПУТЬ: %s", err.Error(), op)
	}

	return nil
}
