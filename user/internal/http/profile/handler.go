package profile

import (
	"context"
	"net/http"

	"messenger/shared/middleware"
	"messenger/user/internal/domain"

	"github.com/go-chi/chi/v5"
)

type UserProfileUseCase interface {
	UserProfile(ctx context.Context) (domain.User, error)
}

type HandlerConfig struct {
	UserProfileUseCase UserProfileUseCase
}

type Handler struct {
	userProfileUseCase UserProfileUseCase
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func New(cfg HandlerConfig) *Handler {
	return &Handler{
		userProfileUseCase: cfg.UserProfileUseCase,
	}
}

func (h *Handler) Handlers() http.Handler {
	rg := chi.NewRouter()
	rg.Use(middleware.ParseToken("test"))
	rg.Group(func(r chi.Router) {
		r.Get("/", h.UserProfile)
	})

	return rg
}
