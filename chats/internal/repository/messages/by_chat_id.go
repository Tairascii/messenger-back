package messages

import (
	"context"

	"messenger/chats/internal/domain"

	"github.com/google/uuid"
)

func (r *repository) ByChatID(ctx context.Context, chatID uuid.UUID) ([]domain.Message, error) {
	var rows []row
	if err := r.db.SelectContext(ctx, &rows, byChatIDSQL, chatID); err != nil {
		return nil, err
	}

	messages := make([]domain.Message, len(rows))
	for i, row := range rows {
		messages[i] = row.toDomain()
	}
	return messages, nil
}

const byChatIDSQL = `
	select id, text, is_edited, created_at, sender_id, chat_id
	from messages
	where chat_id = $1;
`
