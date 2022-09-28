package router

import (
	"github.com/go-chi/chi/v5"
)

func (h Handler) routes() {

	h.Chi.Route("/api/v1", func(r chi.Router) {

	})
}
