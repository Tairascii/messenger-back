package chatsparticipants

import (
	"context"

	"github.com/google/uuid"
)

func (r *repository) Users(ctx context.Context, chatID uuid.UUID) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := r.db.SelectContext(ctx, &ids, usersSQL, chatID); err != nil {
		return nil, err
	}

	return ids, nil
}

const usersSQL = `
		select user_id
		from chats_participants
		where chat_id = $1;
	`
