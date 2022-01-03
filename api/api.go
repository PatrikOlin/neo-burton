package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"

	"github.com/PatrikOlin/neo-burton/handlers"
	m "github.com/PatrikOlin/neo-burton/middleware"
)

func GetRouter(log *zap.Logger) *chi.Mux {
	r := chi.NewRouter()

	if log != nil {
		r.Use(
			middleware.RequestID,
			m.SetLogger(log),
		)
	}

	r.Post("/user", handlers.AddUser)
	r.Post("/signin/{id}", handlers.SignIn)

	return r
}
