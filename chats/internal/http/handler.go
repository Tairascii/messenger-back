package http

import (
	"messenger/chats/internal/http/chats"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type Handler struct {
	chatsHandlers *chats.Handler
}

type Config struct {
	ChatsHandlers *chats.Handler
}

func New(cfg *Config) *Handler {
	return &Handler{
		chatsHandlers: cfg.ChatsHandlers,
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
			v1.Mount("/chats", h.chatsHandlers.Handlers())
		})
	})

	return r
}
