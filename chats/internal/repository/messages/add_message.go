package messages

import (
	"context"
	"database/sql"

	"messenger/chats/internal/domain"
)

func (r *repository) AddMessage(ctx context.Context, message domain.Message) (id int64, err error) {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
	})
	if err != nil {
		return 0, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()
	err = r.db.QueryRowContext(ctx, addMessageSQL,
		message.Text,
		message.IsEdited,
		message.SenderID,
		message.ChatID).Scan(&id)
	if err != nil {
		return 0, err
	}

	_, err = r.db.ExecContext(ctx, updateChatLastMessage, id, message.ChatID)
	if err != nil {
		return 0, err
	}

	return id, nil
}

const addMessageSQL = `
		insert into messages (text, is_edited, sender_id, chat_id)
		values ($1, $2, $3, $4) returning id;
		`

const updateChatLastMessage = `
		update chats
		set last_message_id = $1
		where id = $2;
		`
