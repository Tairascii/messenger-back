package signup

import (
	"context"
	"messenger/user/internal/domain"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (u *UseCase) SignUp(ctx context.Context, req Request) error {
	passwordHash, err := hashPassword(req.Password)
	if err != nil {
		return err
	}

	user := domain.User{
		ID: uuid.New(),
		Email: req.Email,
		Password: passwordHash,
	}
	err = u.userRepo.Create(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func hashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}
