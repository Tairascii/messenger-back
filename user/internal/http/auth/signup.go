package auth

import (
	"encoding/json"
	"net/http"

	"messenger/shared/responsewriter"
	"messenger/user/internal/usecase/auth/signup"
)

// SignUp
//
//	@Summary		Sign up user
//	@Description	sign up user
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		signUpPayload	true	"payload"
//	@Success		201		"ok"
//	@Failure		400		{object}	ErrorResponse
//	@Failure		500		{object}	ErrorResponse
//	@Router			/api/v1/auth/sign-up [post]
func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var payload signUpPayload
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil {
		responsewriter.ErrorResponseWriter(w, err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	req := signup.Request{
		Email:    payload.Email,
		Password: payload.Password,
	}
	err := h.singUpUseCase.SignUp(ctx, req)
	if err != nil {
		responsewriter.ErrorResponseWriter(w, err, http.StatusBadRequest)
		return
	}

	responsewriter.JSONResponseWriter(w, http.StatusCreated, nil)
}

type signUpPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
