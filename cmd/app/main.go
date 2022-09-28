package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog"
	"littlevincee.com/pokemon-api-gateway/internal/pkg/server"
)

func main() {
	logger := httplog.NewLogger("Pokemon API Gateway", httplog.Options{
		JSON:    true,
		Concise: true,
	})

	s := server.New()

	r := chi.NewRouter()
	r.Use(httplog.RequestLogger(logger))
	r.Use(middleware.Heartbeat("/ping"))

	s.Start(":8001", r)
}
