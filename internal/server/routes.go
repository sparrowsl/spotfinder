package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.CleanPath)
	router.Use(middleware.StripSlashes)
	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		message := fmt.Sprintf("%s does not exists or not implemented yet!!", r.URL)
		json.NewEncoder(w).Encode(map[string]any{"error": message})
	})

	router.Mount("/v1", v1Routes())

	return router
}

func v1Routes() *chi.Mux {
	v1Routes := chi.NewRouter()

	v1Routes.Get("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]any{"message": "This was fun!!"})
	})

	return v1Routes
}
