package domain

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrNotParticipant = errors.New("user is not chat participant")
)

type Chat struct {
	ID                uuid.UUID
	LastMessage       Message // todo
	LastReadMessageID int64
}
