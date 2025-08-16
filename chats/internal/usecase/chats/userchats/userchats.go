package userchats

import (
	"context"

	"messenger/chats/internal/domain"
	"messenger/shared/contextutil"
)

// todo pagination
func (u *UseCase) UserChats(ctx context.Context) ([]domain.Chat, error) {
	userID, err := contextutil.UserID(ctx)
	if err != nil {
		return nil, err
	}

	chats, err := u.chatsRepo.ByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return chats, nil
}
