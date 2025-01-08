package api

import (
	"authentication-service/api/cmd"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Post("/authentication-service/logar", cmd.Logar)
	r.Post("/authentication-service/create-user", cmd.CreateUser)

	return r
}
