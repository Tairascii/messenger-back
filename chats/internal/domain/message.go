package domain

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID        int64
	Text      string
	IsEdited  bool
	CreatedAt time.Time
	SenderID  uuid.UUID
	ChatID    uuid.UUID
}
