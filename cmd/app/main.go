package main

import (
	"github.com/go-chi/chi/v5"
	"littlevincee.com/pokemon-api-gateway/internal/pkg/router"
	"littlevincee.com/pokemon-api-gateway/internal/pkg/server"
)

func main() {
	r := &router.Handler{
		Chi: chi.NewRouter(),
	}
	r.Setup()

	s := server.New()
	s.Start("0.0.0.0:8080", r.Chi)
}
