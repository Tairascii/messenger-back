package chats

import (
	"context"
	"time"

	"messenger/chats/internal/domain"

	"github.com/google/uuid"
)

func (r *repository) ByUserID(ctx context.Context, userID uuid.UUID) ([]domain.Chat, error) {
	var rows []chatRow
	if err := r.db.SelectContext(ctx, &rows, byUserIDSQL, userID); err != nil {
		return nil, err
	}

	chats := make([]domain.Chat, len(rows))
	for i, row := range rows {
		chats[i] = row.toDomain()
	}
	return chats, nil
}

type chatRow struct {
	ID                uuid.UUID  `db:"id"`
	Picture           *string    `db:"picture"`
	Title             string     `db:"title"`
	LastReadMessageID *int64     `db:"last_read_message_id"`
	Text              *string    `db:"text"`
	CreatedAt         *time.Time `db:"created_at"`
	SenderID          *uuid.UUID `db:"sender_id"`
}

func (row *chatRow) toDomain() domain.Chat {
	return domain.Chat{
		ID:      row.ID,
		Picture: row.Picture,
		Title:   row.Title,
		LastMessage: domain.LastMessage{
			Text:      row.Text,
			CreatedAt: row.CreatedAt,
			SenderID:  row.SenderID,
		},
		LastReadMessageID: row.LastReadMessageID,
	}
}

const byUserIDSQL = `
		select ch.id, ch.picture, ch.title, ch.last_read_message_id, m.text, m.created_at, m.sender_id
		from chats ch 
		join chats_participants chp on ch.id = chp.chat_id
		left join messages m on ch.last_read_message_id = m.id
		where chp.user_id = $1;
	`
