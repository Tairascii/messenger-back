package contextutil

import (
	"context"
	"errors"

	"messenger/shared/middleware"

	"github.com/google/uuid"
)

var ErrInvalidUserID = errors.New("invalid user id")

func UserID(ctx context.Context) (uuid.UUID, error) {
	userIDRaw, ok := ctx.Value(middleware.UserIDKey).(string)
	if !ok {
		return uuid.Nil, ErrInvalidUserID
	}

	userID, err := uuid.Parse(userIDRaw)
	if err != nil {
		return uuid.Nil, ErrInvalidUserID
	}

	return userID, nil
}
