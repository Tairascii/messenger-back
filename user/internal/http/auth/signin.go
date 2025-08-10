package auth

import (
	"encoding/json"
	"net/http"

	"messenger/shared/responsewriter"
	"messenger/user/internal/domain"
	"messenger/user/internal/usecase/auth/signin"
)

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var payload signUpPayload
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil {
		responsewriter.ErrorResponseWriter(w, err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	req := signin.Request{
		Email:    payload.Email,
		Password: payload.Password,
	}
	tokens, err := h.signInUseCase.SignIn(ctx, req)
	if err != nil {
		responsewriter.ErrorResponseWriter(w, err, http.StatusBadRequest)
		return
	}

	responsewriter.JSONResponseWriter(w, http.StatusCreated, mapDomainToResponse(tokens))
}

type signInPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type singInResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func mapDomainToResponse(resp domain.AccessTokens) singInResponse {
	return singInResponse{
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
	}
}
