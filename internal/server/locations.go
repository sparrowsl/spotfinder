package server

import (
	"context"
	"net/http"
	"spotfinder/internal/database"
	"spotfinder/internal/models"
	"time"

	z "github.com/Oudwins/zog"
	"github.com/Oudwins/zog/zhttp"
)

func (app *Application) getLocations(writer http.ResponseWriter, request *http.Request) {
	category := request.URL.Query().Get("category")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if category != "" {
		locations, err := app.db.GetLocationsByCategory(ctx, &category)
		if err != nil {
			jsonResp(writer, http.StatusInternalServerError, map[string]any{"error": err.Error()})
			return
		}

		jsonResp(writer, http.StatusOK, map[string]any{"locations": locations})
	} else {
		locations, err := app.db.GetLocations(ctx)
		if err != nil {
			jsonResp(writer, http.StatusInternalServerError, map[string]any{"error": err.Error()})
			return
		}

		jsonResp(writer, http.StatusOK, map[string]any{"locations": locations})
	}
}

func (app *Application) addLocation(writer http.ResponseWriter, request *http.Request) {
	input := models.Location{}

	locationSchema := z.Struct(z.Schema{
		"latitude":    z.Float().Required(),
		"longitude":   z.Float().Required(),
		"address":     z.String().Trim().Min(6).Required(),
		"category":    z.String().Trim(),
		"description": z.String().Trim(),
	})
	if err := locationSchema.Parse(zhttp.Request(request), &input); err != nil {
		jsonResp(writer, http.StatusBadRequest, map[string]any{"error": err})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	location, err := app.db.AddLocations(ctx, database.AddLocationsParams{
		Address:     input.Address,
		Latitude:    input.Latitude,
		Longitude:   input.Longitude,
		Category:    &input.Category,
		Description: &input.Description,
	})
	if err != nil {
		jsonResp(writer, http.StatusInternalServerError, map[string]any{"error": err.Error()})
		return
	}

	jsonResp(writer, http.StatusOK, map[string]any{"location": location})
}
