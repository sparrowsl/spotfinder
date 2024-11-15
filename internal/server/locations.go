package server

import (
	"context"
	"net/http"
	"spotfinder/internal/models"
	"time"

	z "github.com/Oudwins/zog"
	"github.com/Oudwins/zog/zhttp"
)

func (app *Application) getLocations(writer http.ResponseWriter, request *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	locations, err := app.db.GetAllLocations(ctx)
	if err != nil {
		jsonResp(writer, http.StatusInternalServerError, map[string]any{"error": err.Error()})
		return
	}

	jsonResp(writer, http.StatusOK, map[string]any{"locations": locations})
}

func (app *Application) addLocation(writer http.ResponseWriter, request *http.Request) {
	input := models.Location{}

	locationSchema := z.Struct(z.Schema{
		"latitude":    z.Float().Required(),
		"longitude":   z.Float().Required(),
		"address":     z.String().Trim().Min(6).Required(),
		"category":    z.String().Trim().Min(4).Required(),
		"description": z.String().Trim(),
	})
	if err := locationSchema.Parse(zhttp.Request(request), &input); err != nil {
		jsonResp(writer, http.StatusBadRequest, map[string]any{"error": err})
		return
	}

	jsonResp(writer, http.StatusOK, map[string]any{"location": input})
}
