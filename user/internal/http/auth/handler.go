package auth

import (
	"context"
	"net/http"

	"messenger/user/internal/domain"
	"messenger/user/internal/usecase/auth/signin"
	"messenger/user/internal/usecase/auth/signup"

	"github.com/go-chi/chi/v5"
)

type SignInUseCase interface {
	SignIn(ctx context.Context, req signin.Request) (domain.AccessTokens, error)
}

type SignUpUseCase interface {
	SignUp(ctx context.Context, req signup.Request) error
}

type HandlerConfig struct {
	SignInUseCase SignInUseCase
	SignUpUseCase SignUpUseCase
}

type Handler struct {
	signInUseCase SignInUseCase
	singUpUseCase SignUpUseCase
}

func New(cfg HandlerConfig) *Handler {
	return &Handler{
		signInUseCase: cfg.SignInUseCase,
		singUpUseCase: cfg.SignUpUseCase,
	}
}

func (h *Handler) Handlers() http.Handler {
	rg := chi.NewRouter()
	rg.Group(func(r chi.Router) {
		r.Post("/sign-in", h.SignIn)
		r.Post("/sign-up", h.SignUp)
	})

	return rg
}
