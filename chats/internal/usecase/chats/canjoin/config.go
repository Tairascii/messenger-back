package canjoin

import (
	"context"

	"github.com/google/uuid"
)


type ChatsParticipantsRepo interface {
	IsParticipant(ctx context.Context, userID, chatID uuid.UUID) (bool, error)
}

type Config struct {
	ChatsParticipantsRepo ChatsParticipantsRepo
}