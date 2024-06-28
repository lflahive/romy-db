package routes

import (
	"net/http"
	"runtime"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(router *chi.Mux) {
	router.Route("/api", func(r chi.Router) {
		r.Get("/health", health)
		mapCollectionRoutes(r)
		mapPartitionRoutes(r)
	})
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(runtime.Version()))
}
