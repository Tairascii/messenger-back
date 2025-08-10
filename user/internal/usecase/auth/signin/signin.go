package signin

import (
	"context"
	"messenger/user/internal/domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func (u *UseCase) SignIn(ctx context.Context, req Request) (domain.AccessTokens, error) {
	user, err := u.userRepo.ByEmail(ctx, req.Email)
	if err != nil {
		return domain.AccessTokens{}, nil
	}

	if err = checkPassword(req.Password, user.Password); err != nil {
		return domain.AccessTokens{}, domain.ErrUserNotFound
	}

	id := user.ID.String()
	accessToken, err := generateToken(req.Email, id, "test", time.Now().Add(24*time.Hour).Unix())
	if err != nil {
		return domain.AccessTokens{}, err
	}
	refreshToken, err := generateToken(req.Email, id, "test", time.Now().Add(7*24*time.Hour).Unix())
	if err != nil {
		return domain.AccessTokens{}, err
	}
	
	return domain.AccessTokens{
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func checkPassword(password, passwordHash string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
}

func generateToken(email, id, secret string, exp int64) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"id":    id,
		"iat":   time.Now().Unix(),
		"exp":   exp,
	}).SignedString([]byte(secret))
}
