package addmessage

import (
	"context"

	"messenger/chats/internal/domain"
	"messenger/shared/contextutil"
)

func (u *UseCase) AddMessage(ctx context.Context, message domain.Message) error {
	userID, err := contextutil.UserID(ctx)
	if err != nil {
		return err
	}

	message.SenderID = userID
	isParticipant, err := u.chatsParticipantsRepo.IsParticipant(ctx, userID, message.ChatID)
	if err != nil {
		return err
	}

	if !isParticipant {
		return domain.ErrNotParticipant
	}

	err = u.AddMessage(ctx, message)
	if err != nil {
		return err
	}

	return nil
}
