package userchats

import (
	"context"
	"messenger/chats/internal/domain"

	"github.com/google/uuid"
)

type ChatsRepo interface {
	ByUserID(ctx context.Context, userID uuid.UUID) ([]domain.Chat, error)
}

type Config struct {
	ChatsRepo ChatsRepo
}
