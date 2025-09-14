package router

import (
	"github.com/AriSu2904/go-auth/internal/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(authHandler handler.AuthHandler, userHandler handler.UserHandler) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", authHandler.Register)
	})

	r.Route("/user", func(r chi.Router) {
		r.Get("/", userHandler.FindByQuery)
	})

	return r
}
