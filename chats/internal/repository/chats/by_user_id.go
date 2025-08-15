package chats

import (
	"context"

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
	ID                uuid.UUID `db:"id"`
	LastReadMessageID int64     `db:"last_read_message_id"`
}

func (row *chatRow) toDomain() domain.Chat {
	return domain.Chat{
		ID:                row.ID,
		LastReadMessageID: row.LastReadMessageID,
	}
}

const byUserIDSQL = `
		select id, last_read_message_id
		from chats ch 
		join chats_participants chp on ch.id = chp.id
		where user_id = $1;
	`
