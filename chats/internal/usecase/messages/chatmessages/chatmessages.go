package chatmessages

import (
	"context"

	"messenger/chats/internal/domain"
	"messenger/shared/contextutil"

	"github.com/google/uuid"
)

func (u *UseCase) ChatMessages(ctx context.Context, chatID uuid.UUID) ([]domain.Message, error) {
	userID, err := contextutil.UserID(ctx)
	if err != nil {
		return nil, err
	}

	isParticipant, err := u.chatsParticipantsRepo.IsParticipant(ctx, userID, chatID)
	if err != nil {
		return nil, err
	}

	if !isParticipant {
		return nil, domain.ErrNotParticipant
	}

	messages, err := u.messageRepo.ByChatID(ctx, chatID)
	if err != nil {
		return nil, err
	}

	return messages, nil
}
