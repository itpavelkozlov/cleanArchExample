package server

import (
	"cleanArch/internal/resolver"
	"github.com/go-chi/chi/v5"
)

func NewRouter(resolver *resolver.Resolver) *chi.Mux {
	r := chi.NewRouter()
	r.Post("/createUser", resolver.CreateUser)
	//r.Get("/getUser", resolver.GetUser)
	return r
}
