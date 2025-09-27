package router

import (
	"github.com/AriSu2904/go-auth/internal/handler"
	"github.com/AriSu2904/go-auth/internal/middleware"
	"github.com/go-chi/chi/v5"
)

func NewRouter(authHandler handler.AuthHandler, userHandler handler.UserHandler) *chi.Mux {
	r := chi.NewRouter()

	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", authHandler.Register)
		r.With(middleware.HeaderValidator).Post("/login", authHandler.Login)
	})

	r.Route("/user", func(r chi.Router) {
		r.Get("/", userHandler.FindByQuery)
	})

	return r
}
