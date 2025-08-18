package addmessage

import (
	"context"

	"messenger/chats/internal/domain"
	"messenger/shared/contextutil"
)

func (u *UseCase) AddMessage(ctx context.Context, message domain.Message) (domain.Message, error) {
	userID, err := contextutil.UserID(ctx)
	if err != nil {
		return domain.Message{}, err
	}

	message.SenderID = userID
	isParticipant, err := u.chatsParticipantsRepo.IsParticipant(ctx, userID, message.ChatID)
	if err != nil {
		return domain.Message{}, err
	}

	if !isParticipant {
		return domain.Message{}, domain.ErrNotParticipant
	}

	id, err := u.messageRepo.AddMessage(ctx, message)
	if err != nil {
		return domain.Message{}, err
	}

	return domain.Message{
		ID:        id,
		Text:      message.Text,
		IsEdited:  message.IsEdited,
		CreatedAt: message.CreatedAt,
		SenderID:  message.SenderID,
		ChatID:    message.ChatID,
	}, nil
}
