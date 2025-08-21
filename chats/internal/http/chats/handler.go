package chats

import (
	"context"
	"net/http"
	"sync"

	"messenger/chats/internal/domain"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type UserChatsUseCase interface {
	UserChats(ctx context.Context) ([]domain.Chat, error)
}

type DeleteChatUseCase interface {
	DeleteByID(ctx context.Context, id uuid.UUID) error
}

type AddMessageUseCase interface {
	AddMessage(ctx context.Context, message domain.Message) (domain.Message, error)
}

type CanJoinUseCase interface {
	CanJoin(ctx context.Context, chatID uuid.UUID) (uuid.UUID, error)
}

type CreateChatUseCase interface {
	CreateChat(ctx context.Context, user2 uuid.UUID) (uuid.UUID, error)
}

type HandlerConfig struct {
	UserChatsUseCase  UserChatsUseCase
	DeleteChatUseCase DeleteChatUseCase
	AddMessageUseCase AddMessageUseCase
	CanJoinUseCase    CanJoinUseCase
	CreateChatUseCase CreateChatUseCase
}

type Handler struct {
	userChatsUseCase  UserChatsUseCase
	deleteChatUseCase DeleteChatUseCase
	addMessageUseCase AddMessageUseCase
	canJoinUseCase    CanJoinUseCase
	createChatUseCase CreateChatUseCase
	chatConnections   map[uuid.UUID]map[uuid.UUID]*websocket.Conn
	mu                *sync.Mutex
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func New(cfg HandlerConfig) *Handler {
	return &Handler{
		userChatsUseCase:  cfg.UserChatsUseCase,
		deleteChatUseCase: cfg.DeleteChatUseCase,
		addMessageUseCase: cfg.AddMessageUseCase,
		canJoinUseCase:    cfg.CanJoinUseCase,
		createChatUseCase: cfg.CreateChatUseCase,
		chatConnections:   make(map[uuid.UUID]map[uuid.UUID]*websocket.Conn),
		mu:                &sync.Mutex{},
	}
}

func (h *Handler) Handlers() http.Handler {
	rg := chi.NewRouter()
	rg.Group(func(r chi.Router) {
		r.Get("/", h.UserChats)
		r.Post("/", h.CreateChat)
		r.Delete("/{chat_id}", h.DeleteChat)
		r.HandleFunc("/connect/{chat_id}", h.ConnectToChat)
	})

	return rg
}
