package user

import (
	"context"

	"messenger/user/internal/domain"

	"github.com/google/uuid"
)

func (r *repository) ByEmail(ctx context.Context, email string) (domain.User, error) {
	q := `
		select id, email, password
		from users
		where email = $1;
	`
	var row userRow
	if err := r.db.GetContext(ctx, &row, q); err != nil {
		return domain.User{}, err
	}

	return row.toDomain(), nil
}

type userRow struct {
	ID       uuid.UUID `db:"id"`
	Email    string    `db:"email"`
	Password string    `db:"password"`
}

func (row *userRow) toDomain() domain.User {
	return domain.User{
		ID:       row.ID,
		Email:    row.Email,
		Password: row.Password,
	}
}
