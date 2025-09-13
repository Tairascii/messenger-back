package messages

import (
	"context"
	"net/http"

	"messenger/chats/internal/domain"
	"messenger/shared/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type ChatMessagesUseCase interface {
	ChatMessages(ctx context.Context, chatID uuid.UUID) ([]domain.Message, error)
}

type HandlerConfig struct {
	ChatMessagesUseCase ChatMessagesUseCase
}

type Handler struct {
	chatMessagesUseCase ChatMessagesUseCase
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func New(cfg HandlerConfig) *Handler {
	return &Handler{
		chatMessagesUseCase: cfg.ChatMessagesUseCase,
	}
}

func (h *Handler) Handlers() http.Handler {
	rg := chi.NewRouter()
	rg.Use(middleware.ParseToken("test"))
	rg.Group(func(r chi.Router) {
		r.Get("/{chat_id}", h.ChatMessages)
	})

	return rg
}
