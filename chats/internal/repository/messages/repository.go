package messages

import (
	"messenger/chats/internal/domain"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *repository {
	return &repository{
		db: db,
	}
}

type row struct {
	ID        int64     `db:"id"`
	Text      string    `db:"text"`
	IsEdited  bool      `db:"is_edited"`
	CreatedAt time.Time `db:"created_at"`
	SenderID  uuid.UUID `db:"sender_id"`
	ChatID    uuid.UUID `db:"chat_id"`
}

func (r row) toDomain() domain.Message {
	return domain.Message(r)
}
