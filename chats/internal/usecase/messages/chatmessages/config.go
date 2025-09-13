package chatmessages

import (
	"context"

	"messenger/chats/internal/domain"

	"github.com/google/uuid"
)

type MessagesRepo interface {
	ByChatID(ctx context.Context, chatID uuid.UUID) ([]domain.Message, error)
}

type ChatsParticipantsRepo interface {
	IsParticipant(ctx context.Context, userID, chatID uuid.UUID) (bool, error)
}

type Config struct {
	MessagesRepo          MessagesRepo
	ChatsParticipantsRepo ChatsParticipantsRepo
}
