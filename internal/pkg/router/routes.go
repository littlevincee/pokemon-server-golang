package router

import (
	"github.com/go-chi/chi/v5"
)

func (h Handler) routes() {
	//l := logger.New()

	h.chi.Route("/api/v1", func(r chi.Router) {

	})
}
