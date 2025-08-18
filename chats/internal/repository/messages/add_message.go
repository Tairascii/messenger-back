package messages

import (
	"context"

	"messenger/chats/internal/domain"

	"github.com/google/uuid"
)

func (r *repository) AddMessage(ctx context.Context, message domain.Message) (int64, error) {
	var id int64
	params := row{
		Text:     message.Text,
		IsEdited: message.IsEdited,
		SenderID: message.SenderID,
		ChatID:   message.ChatID,
	}
	err := r.db.QueryRowContext(ctx, addMessageSQL, params).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

type row struct {
	Text     string    `db:"text"`
	IsEdited bool      `db:"is_edited"`
	SenderID uuid.UUID `db:"sender_id"`
	ChatID   uuid.UUID `db:"chat_id"`
}

const addMessageSQL = `
		insert into messages (text, is_edited, sender_id, chat_id)
		values (:text, :is_edited, :sender_id, :chat_id) returning id;
		`
