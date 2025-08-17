package messages

import (
	"context"

	"messenger/chats/internal/domain"

	"github.com/google/uuid"
)

func (r *repository) AddMessage(ctx context.Context, message domain.Message) error {
	params := row{
		Text:     message.Text,
		IsEdited: message.IsEdited,
		SenderID: message.SenderID,
		ChatID:   message.ChatID,
	}
	_, err := r.db.ExecContext(ctx, addMessageSQL, params)
	if err != nil {
		return err
	}

	return nil
}

type row struct {
	Text     string    `db:"text"`
	IsEdited bool      `db:"is_edited"`
	SenderID uuid.UUID `db:"sender_id"`
	ChatID   uuid.UUID `db:"chat_id"`
}

const addMessageSQL = `
		insert into messages (text, is_edited, sender_id, chat_id)
		values (:text, :is_edited, :sender_id, :chat_id)
		`
