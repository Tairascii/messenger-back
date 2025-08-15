package domain

import "github.com/google/uuid"

type Chat struct {
	ID                uuid.UUID
	LastMessage       Message // todo
	LastReadMessageID int64
}
