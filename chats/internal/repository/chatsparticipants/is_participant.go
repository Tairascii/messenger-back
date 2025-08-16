package chatsparticipants

import (
	"context"

	"github.com/google/uuid"
)

func (r *repository) IsParticipant(ctx context.Context, userID, chatID uuid.UUID) (bool, error) {
	var exists bool
	if err := r.db.GetContext(ctx, &exists, isParticipantSQL); err != nil {
		return false, err
	}

	return exists, nil
}

const isParticipantSQL = `
		select exists(
			select 1
			from chats_participants
			where chat_id = $1 and user_id = $2
		);
	`
