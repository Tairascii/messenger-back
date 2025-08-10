package domain

import "errors"

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
)

type AccessTokens struct {
	AccessToken  string
	RefreshToken string
}
