package canjoin

import (
	"context"

	"messenger/chats/internal/domain"
	"messenger/shared/contextutil"

	"github.com/google/uuid"
)

func (u *UseCase) CanJoin(ctx context.Context, chatID uuid.UUID) (uuid.UUID, error) {
	userID, err := contextutil.UserID(ctx)
	if err != nil {
		return uuid.Nil, err
	}

	can, err := u.chatsParticipantsRepo.IsParticipant(ctx, userID, chatID)
	if err != nil {
		return uuid.Nil, err
	}

	if !can {
		return uuid.Nil, domain.ErrNotParticipant
	}
	return userID, nil
}
