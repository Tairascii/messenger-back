package createchat

import (
	"context"

	"messenger/shared/contextutil"

	"github.com/google/uuid"
)

func (u *UseCase) CreateChat(ctx context.Context, user2 uuid.UUID) (uuid.UUID, error) {
	userID, err := contextutil.UserID(ctx)
	if err != nil {
		return uuid.Nil, err
	}

	chatID := uuid.New()
	err = u.chatsRepo.Create(ctx, chatID, userID, user2)
	if err != nil {
		return uuid.Nil, err
	}

	return chatID, nil
}
