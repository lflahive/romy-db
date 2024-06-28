package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lflahive/romy-db/internal/collection"
)

func mapCollectionRoutes(r chi.Router) {
	r.Post("/collection/{name}", createCollection)
	r.Get("/collection/{name}", getCollection)
}

func createCollection(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	if err := collection.Create(name); err != nil {
		writeError(w, err, http.StatusBadRequest)
		return
	}
}

func getCollection(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	info, err := collection.Get(name)
	if err != nil {
		writeError(w, err, http.StatusBadRequest)
		return
	}

	writeResponse(w, info)
}
