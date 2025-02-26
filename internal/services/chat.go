package services

import (
	"context"
	"fmt"

	"github.com/gene-qxsi/Flexive/internal/domain"
	"github.com/gene-qxsi/Flexive/internal/repository/models"
)

type ChatRepo interface {
	GetChats(ctx context.Context) ([]models.Chat, error)
	CreateChat(ctx context.Context, chat *models.Chat) (*models.Chat, error)
	GetMessagesByChatID(ctx context.Context, chatID, count int) ([]models.Message, error)
	CreateMessage(ctx context.Context, chatID int, message *models.Message) error
}

type ChatService struct {
	repo ChatRepo
}

func NewChatService(chatRepo ChatRepo) *ChatService {
	return &ChatService{repo: chatRepo}
}

func (s *ChatService) GetChats(ctx context.Context) ([]domain.Chat, error) {
	// const op = "internal/services/chat.go/GetChats()"

	models, err := s.repo.GetChats(ctx)
	if err != nil {
		return nil, err
	}

	var chats []domain.Chat
	for _, chat := range models {
		chats = append(chats, *domain.ToDomainChat(&chat))
	}

	return chats, nil
}

func (s *ChatService) CreateChat(ctx context.Context, chat domain.Chat) (*domain.Chat, error) {
	// const op = "internal/services/chat.go/CreateChat()"

	req := domain.ToORMChat(&chat)

	model, err := s.repo.CreateChat(ctx, req)
	if err != nil {
		return nil, err
	}
	fmt.Println("НОРМОООООЛЬНО 2222: 2", chat)
	return domain.ToDomainChat(model), nil
}

func (s *ChatService) GetMessagesByChatID(ctx context.Context, chatID, limit int) ([]domain.Message, error) {
	// const op = "internal/services/chat.go/GetMessagesByChatID()"

	models, err := s.repo.GetMessagesByChatID(ctx, chatID, limit)
	if err != nil {
		return nil, err
	}

	var messages []domain.Message
	for _, chat := range models {
		messages = append(messages, *domain.ToDomainMessage(&chat))
	}

	return messages, nil
}

func (s *ChatService) CreateMessage(ctx context.Context, chatID int, message domain.Message) error {
	// const op = "internal/services/chat.go/v()"

	ormMessage := domain.ToORMMessage(&message)

	err := s.repo.CreateMessage(ctx, chatID, ormMessage)
	if err != nil {
		return err
	}

	return nil
}
