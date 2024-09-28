package handlers

import (
	"github.com/MouraJr/firstgoapi/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) *chi.Mux {
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.router) {
		router.Use(middleware.Authorization)

		router.Get("/notes", GetNotes)
	})
}
