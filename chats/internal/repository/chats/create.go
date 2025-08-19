package chats

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

func (r *repository) Create(ctx context.Context, chatID, user1, user2 uuid.UUID) (err error) {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
	})
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()
	params := createChatParams{
		ID: chatID,
	}
	_, err = r.db.NamedExecContext(ctx, createChatSQL, params)
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, addParticipantsSQL, chatID, user1)
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, addParticipantsSQL, chatID, user2)
	if err != nil {
		return err
	}

	return nil
}

type createChatParams struct {
	ID uuid.UUID `db:"id"`
}

const createChatSQL = `
	insert into chats (id)
	values (:id);
`

const addParticipantsSQL = `
	insert into chats_participants (chat_id, user_id)
	values ($1, $2);
`
