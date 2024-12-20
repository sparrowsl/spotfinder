package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *Application) routes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.CleanPath)
	router.Use(middleware.StripSlashes)
	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		message := fmt.Sprintf("%s with %s method does not exists!!", r.URL, r.Method)
		jsonResp(w, http.StatusNotFound, map[string]any{"error": message})
	})

	router.Mount("/v1", app.v1())

	return router
}

func (app *Application) v1() *chi.Mux {
	v1 := chi.NewRouter()

	v1.Get("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]any{"message": "This was fun!!"})
	})

	v1.Get("/locations", app.getLocations)
	v1.Post("/locations", app.addLocation)

	return v1
}
