package deletechat

import (
	"context"

	"messenger/chats/internal/domain"
	"messenger/shared/contextutil"

	"github.com/google/uuid"
)

func (u *UseCase) DeleteByID(ctx context.Context, id uuid.UUID) error {
	userID, err := contextutil.UserID(ctx)
	if err != nil {
		return err
	}

	isParticipant, err := u.chatsParticipantsRepo.IsParticipant(ctx, userID, id)
	if err != nil {
		return err
	}

	if !isParticipant {
		return domain.ErrNotParticipant
	}

	err = u.chatsParticipantsRepo.DeleteByID(ctx, id, userID)
	if err != nil {
		return err
	}

	return nil
}
