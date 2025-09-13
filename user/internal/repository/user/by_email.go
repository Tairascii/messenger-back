package user

import (
	"context"

	"messenger/user/internal/domain"
)

func (r *repository) ByEmail(ctx context.Context, email string) (domain.User, error) {
	var row userRow
	if err := r.db.GetContext(ctx, &row, byEmailSQL, email); err != nil {
		return domain.User{}, err
	}

	return row.toDomain(), nil
}

const byEmailSQL = `
		select id, email, password
		from users
		where email = $1;
	`
