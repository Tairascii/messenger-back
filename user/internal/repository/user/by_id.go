package user

import (
	"context"

	"messenger/user/internal/domain"

	"github.com/google/uuid"
)

func (r *repository) ByID(ctx context.Context, id uuid.UUID) (domain.User, error) {
	var row userRow
	if err := r.db.GetContext(ctx, &row, byIDSQL, id); err != nil {
		return domain.User{}, err
	}

	return row.toDomain(), nil
}

const byIDSQL = `
		select id, email, password
		from users
		where id = $1;
	`
