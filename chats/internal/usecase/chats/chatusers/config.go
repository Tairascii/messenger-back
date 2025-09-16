package chatusers

import (
	"context"

	"github.com/google/uuid"
)


type ChatsParticipantsRepo interface {
	Users(ctx context.Context, chatID uuid.UUID) ([]uuid.UUID, error)
}

type Config struct {
	ChatsParticipantsRepo ChatsParticipantsRepo
}
