package userprofile

import (
	"context"

	"messenger/shared/contextutil"
	"messenger/user/internal/domain"
)

func (u *UseCase) UserProfile(ctx context.Context) (domain.User, error) {
	userID, err := contextutil.UserID(ctx)
	if err != nil {
		return domain.User{}, err
	}

	user, err := u.userRepo.ByID(ctx, userID)
	if err != nil {
		return domain.User{}, domain.ErrUserNotFound
	}

	return user, nil
}
