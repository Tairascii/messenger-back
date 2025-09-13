package http

import (
	"messenger/chats/internal/http/chats"
	"messenger/chats/internal/http/messages"
	"messenger/shared/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type Handler struct {
	chatsHandlers    *chats.Handler
	messagesHandlers *messages.Handler
}

type Config struct {
	ChatsHandlers    *chats.Handler
	MessagesHandlers *messages.Handler
}

func New(cfg *Config) *Handler {
	return &Handler{
		chatsHandlers:    cfg.ChatsHandlers,
		messagesHandlers: cfg.MessagesHandlers,
	}
}

func (h *Handler) InitHandlers() *chi.Mux {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}))
	r.Use(middleware.ParseToken("test"))
	r.Route("/api", func(api chi.Router) {
		api.Route("/v1", func(v1 chi.Router) {
			v1.Mount("/chats", h.chatsHandlers.Handlers())
			v1.Mount("/messages", h.messagesHandlers.Handlers())
		})
	})

	return r
}
