package userprofile

import (
	"context"

	"messenger/user/internal/domain"

	"github.com/google/uuid"
)

type UserRepo interface {
	ByID(ctx context.Context, id uuid.UUID) (domain.User, error)
}

type Config struct {
	UserRepo UserRepo
}
