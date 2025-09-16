package chatusers

import (
	"context"

	"github.com/google/uuid"
)

func (u *UseCase) ChatUsers(ctx context.Context, chatID uuid.UUID) ([]uuid.UUID, error) {
	ids, err := u.chatsParticipantsRepo.Users(ctx, chatID)
	if err != nil {
		return nil, err
	}

	return ids, nil
}