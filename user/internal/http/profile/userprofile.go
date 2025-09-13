package profile

import (
	"net/http"

	"messenger/shared/responsewriter"
	"messenger/user/internal/domain"

	"github.com/google/uuid"
)

func (h *Handler) UserProfile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, err := h.userProfileUseCase.UserProfile(ctx)
	if err != nil {
		responsewriter.ErrorResponseWriter(w, err, http.StatusBadRequest)
		return
	}

	responsewriter.JSONResponseWriter(w, http.StatusOK, mapDomainToResponse(user))
}

type userProfileResponse struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
}

func mapDomainToResponse(resp domain.User) userProfileResponse {
	return userProfileResponse{
		ID:    resp.ID,
		Email: resp.Email,
	}
}
