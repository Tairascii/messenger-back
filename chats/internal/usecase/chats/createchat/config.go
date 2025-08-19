package createchat

import (
	"context"

	"github.com/google/uuid"
)

type ChatsRepo interface {
	Create(ctx context.Context, chatID, user1, user2 uuid.UUID) (err error)
}

type Config struct {
	ChatsRepo ChatsRepo
}
