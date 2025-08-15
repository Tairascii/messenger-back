package userchats

import (
	"context"
	"messenger/chats/internal/domain"

	"github.com/google/uuid"
)

// todo pagination
func (u *UseCase) UserChats(ctx context.Context, userID uuid.UUID) ([]domain.Chat, error) {
	chats, err := u.chatsRepo.ByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return chats, nil
}
