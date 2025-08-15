package chats

import (
	"context"
	"net/http"

	"messenger/chats/internal/domain"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type UserChatsUseCase interface {
	UserChats(ctx context.Context, userID uuid.UUID) ([]domain.Chat, error)
}

type HandlerConfig struct {
	UserChatsUseCase UserChatsUseCase
}

type Handler struct {
	userChatsUseCase UserChatsUseCase
}

func New(cfg HandlerConfig) *Handler {
	return &Handler{
		userChatsUseCase: cfg.UserChatsUseCase,
	}
}

func (h *Handler) Handlers() http.Handler {
	rg := chi.NewRouter()
	rg.Group(func(r chi.Router) {
		r.Get("/{user_id}", h.UserChats)
	})

	return rg
}
