package deletechat

import (
	"context"

	"github.com/google/uuid"
)

type ChatsParticipantsRepo interface {
	IsParticipant(ctx context.Context, userID, chatID uuid.UUID) (bool, error)
	DeleteByID(ctx context.Context, id, userID uuid.UUID) error
}

type Config struct {
	ChatsParticipantsRepo ChatsParticipantsRepo
}
