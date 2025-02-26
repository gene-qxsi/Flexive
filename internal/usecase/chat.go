package usecase

import (
	"context"

	"github.com/gene-qxsi/Flexive/internal/domain"
)

type ChatService interface {
	GetChats(ctx context.Context) ([]domain.Chat, error)
	CreateChat(ctx context.Context, chat domain.Chat) (*domain.Chat, error)
	GetMessagesByChatID(ctx context.Context, chatID, count int) ([]domain.Message, error)
	CreateMessage(ctx context.Context, chatID int, message domain.Message) error
}

type ChatUsecase struct {
	ChatSrv ChatService
}

func NewChatUsecase(chatSrv ChatService) *ChatUsecase {
	return &ChatUsecase{ChatSrv: chatSrv}
}

func (c *ChatUsecase) LoadLastNMessagesByChatID(ctx context.Context, chatID, count int) ([]domain.Message, error) {
	// const op = "internal/usecase/chat.go/LoadLastNMessagesByChatID()"

	return c.ChatSrv.GetMessagesByChatID(ctx, chatID, count)
}

func (c *ChatUsecase) GetChats(ctx context.Context) ([]domain.Chat, error) {
	// const op = "internal/usecase/chat.go/GetChats()"

	return c.ChatSrv.GetChats(ctx)
}

func (c *ChatUsecase) CreateChat(ctx context.Context, chat domain.Chat) (*domain.Chat, error) {
	// const op = "internal/usecase/chat.go/CreateChat()"
	return c.ChatSrv.CreateChat(ctx, chat)
}

func (c *ChatUsecase) SaveMessage(ctx context.Context, chatID int, message domain.Message) error {
	// const op = "internal/usecase/chat.go/SaveMessage()"

	return c.ChatSrv.CreateMessage(ctx, chatID, message)
}
