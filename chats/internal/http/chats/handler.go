package chats

import (
	"context"
	"net/http"

	"messenger/chats/internal/domain"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type UserChatsUseCase interface {
	UserChats(ctx context.Context) ([]domain.Chat, error)
}

type DeleteChatUseCase interface {
	DeleteByID(ctx context.Context, id uuid.UUID) error
}

type HandlerConfig struct {
	UserChatsUseCase  UserChatsUseCase
	DeleteChatUseCase DeleteChatUseCase
}

type Handler struct {
	userChatsUseCase  UserChatsUseCase
	deleteChatUseCase DeleteChatUseCase
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func New(cfg HandlerConfig) *Handler {
	return &Handler{
		userChatsUseCase:  cfg.UserChatsUseCase,
		deleteChatUseCase: cfg.DeleteChatUseCase,
	}
}

func (h *Handler) Handlers() http.Handler {
	rg := chi.NewRouter()
	rg.Group(func(r chi.Router) {
		r.Get("/", h.UserChats)
		r.Delete("/{chat_id}", h.DeleteChat)
	})

	return rg
}
