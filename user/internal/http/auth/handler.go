package auth

import (
	"context"
	"net/http"

	"messenger/user/internal/domain"
	"messenger/user/internal/usecase/auth/signin"
	"messenger/user/internal/usecase/auth/signup"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
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

func (h *Handler) InitHandlers() *chi.Mux {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}))
	r.Route("/api", func(api chi.Router) {
		api.Route("/v1", func(v1 chi.Router) {
			v1.Mount("/auth", h.authHandlers())
		})
	})
	return r
}

func (h *Handler) authHandlers() http.Handler {
	rg := chi.NewRouter()
	rg.Group(func(r chi.Router) {
		r.Post("/sign-in", h.SignIn)
		r.Post("/sign-up", h.SignUp)
	})

	return rg
}
