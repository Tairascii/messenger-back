package user

import (
	"context"

	"messenger/user/internal/domain"

	"github.com/google/uuid"
)

func (r *repository) Create(ctx context.Context, user domain.User) error {
	params := createParams{
		ID: user.ID,
		Email: user.Email,
		Password: user.Password,
	}
	_, err := r.db.ExecContext(ctx, createSQL, params)
	if err != nil {
		return err
	}

	return nil
}

type createParams struct {
	ID       uuid.UUID `db:"id"`
	Email    string    `db:"email"`
	Password string    `db:"password"`
}

const createSQL = `
	insert into users (id, email, password)
	values (:id, :email, :password);
	`
