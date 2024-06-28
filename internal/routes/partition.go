package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lflahive/romy-db/internal/partition"
)

func mapPartitionRoutes(r chi.Router) {
	r.Post("/partition/{collection}/{name}", createPartition)
	r.Get("/partition/{collection}/{name}", getPartition)
}

func createPartition(w http.ResponseWriter, r *http.Request) {
	collection := chi.URLParam(r, "collection")
	name := chi.URLParam(r, "name")

	if err := partition.Create(collection, name); err != nil {
		writeError(w, err, http.StatusBadRequest)
		return
	}
}

func getPartition(w http.ResponseWriter, r *http.Request) {
	collection := chi.URLParam(r, "collection")
	name := chi.URLParam(r, "name")

	info, err := partition.Get(collection, name)
	if err != nil {
		writeError(w, err, http.StatusBadRequest)
		return
	}

	writeResponse(w, info)
}
