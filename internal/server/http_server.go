package server

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"time"
)

func NewHttpServer(r *chi.Mux) http.Server {
	return http.Server{
		ReadTimeout:       1 * time.Second,
		WriteTimeout:      1 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
		Handler:           r,
	}
}
