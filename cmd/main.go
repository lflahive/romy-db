package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/lflahive/romy-db/internal/config"
	"github.com/lflahive/romy-db/internal/routes"
)

func main() {
	storagePath := flag.String("p", "~/romy-db", "path to store data")
	flag.Parse()

	config.NewConfig(*storagePath)

	fmt.Println(config.Configuration)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))

	routes.RegisterRoutes(r)

	slog.Info("Starting web server...")
	http.ListenAndServe(":8080", r)
}
