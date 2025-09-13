package messages

import (
	"context"

	"messenger/chats/internal/domain"
)

func (r *repository) AddMessage(ctx context.Context, message domain.Message) (int64, error) {
	var id int64
	err := r.db.QueryRowContext(ctx, addMessageSQL,
		message.Text,
		message.IsEdited,
		message.SenderID,
		message.ChatID).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

const addMessageSQL = `
		insert into messages (text, is_edited, sender_id, chat_id)
		values ($1, $2, $3, $4) returning id;
		`
