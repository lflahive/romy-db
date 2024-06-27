package routes

import (
	"net/http"
	"runtime"

	"github.com/go-chi/chi/v5"
	"github.com/lflahive/romy-db/internal/storage"
)

func RegisterRoutes(router *chi.Mux) {
	router.Route("/api", func(r chi.Router) {
		r.Get("/health", health)
		r.Post("/collection/{name}", createCollection)
		r.Get("/collection/{name}", getCollection)
	})
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(runtime.Version()))
}

func createCollection(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	if err := storage.CreateCollection(name); err != nil {
		writeError(w, err, http.StatusBadRequest)
		return
	}
}

func getCollection(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	info, err := storage.GetCollection(name)
	if err != nil {
		writeError(w, err, http.StatusBadRequest)
		return
	}

	writeResponse(w, info)
}
