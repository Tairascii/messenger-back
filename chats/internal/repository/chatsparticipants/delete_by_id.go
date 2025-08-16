package chatsparticipants

import (
	"context"

	"github.com/google/uuid"
)

func (r *repository) DeleteByID(ctx context.Context, id, userID uuid.UUID) error {
	_, err := r.db.ExecContext(ctx, byIDSQL)
	if err != nil {
		return err
	}

	return nil
}

const byIDSQL = `
		delete from chats_participants
		where chat_id = $1 and user_id = $2;
	`
