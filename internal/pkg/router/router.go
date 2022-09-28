package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httplog"
	"time"
)

type Handler struct {
	chi *chi.Mux
}

func (h Handler) Setup() {
	logger := httplog.NewLogger("Pokemon API Gateway", httplog.Options{
		JSON:    true,
		Concise: true,
	})

	h.chi.Use(httplog.RequestLogger(logger))
	h.chi.Use(
		middleware.RequestID,
		middleware.RealIP,
		middleware.Recoverer,
		middleware.Compress(5),
		middleware.Timeout(time.Minute),
		middleware.Heartbeat("/ping"),
		cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           300, // Maximum value not ignored by any of major browsers
		}))

	h.routes()
}
