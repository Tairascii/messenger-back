package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var ErrNotParticipant = errors.New("user is not chat participant")

type Chat struct {
	ID                uuid.UUID
	Picture           *string
	Title             string
	LastMessage       LastMessage
	LastReadMessageID *int64
}

type LastMessage struct {
	Text      *string
	SenderID  *uuid.UUID
	CreatedAt *time.Time
}
