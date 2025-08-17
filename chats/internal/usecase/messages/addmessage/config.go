package addmessage

import (
	"context"

	"messenger/chats/internal/domain"

	"github.com/google/uuid"
)

type MessagesRepo interface {
	AddMessage(ctx context.Context, message domain.Message) error
}

type ChatsParticipantsRepo interface {
	IsParticipant(ctx context.Context, userID, chatID uuid.UUID) (bool, error)
}

type Config struct {
	MessagesRepo          MessagesRepo
	ChatsParticipantsRepo ChatsParticipantsRepo
}
