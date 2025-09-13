package user

import (
	"messenger/user/internal/domain"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *repository {
	return &repository{
		db: db,
	}
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
