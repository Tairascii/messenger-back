package signin

import (
	"context"

	"messenger/user/internal/domain"
)

type UserRepo interface {
	ByEmail(ctx context.Context, email string) (domain.User, error)
}

type Config struct {
	UserRepo UserRepo
}
