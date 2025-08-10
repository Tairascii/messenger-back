package signup

import (
	"context"
	"messenger/user/internal/domain"
)

type UserRepo interface {
	Create(ctx context.Context, user domain.User) error
}

type Config struct {
	UserRepo UserRepo
}
