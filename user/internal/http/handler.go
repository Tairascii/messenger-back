package http

import (
	"messenger/user/internal/http/auth"
	"messenger/user/internal/http/profile"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type Handler struct {
	authHandlers    *auth.Handler
	profileHandlers *profile.Handler
}

type Config struct {
	AuthHandlers    *auth.Handler
	ProfileHandlers *profile.Handler
}

func New(cfg *Config) *Handler {
	return &Handler{
		authHandlers:    cfg.AuthHandlers,
		profileHandlers: cfg.ProfileHandlers,
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
			v1.Mount("/auth", h.authHandlers.Handlers())
			v1.Mount("/profile", h.profileHandlers.Handlers())
		})
	})

	return r
}
