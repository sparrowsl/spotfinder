package server

import (
	"context"
	"net/http"
	"time"
)

func (app *Application) getLocations(writer http.ResponseWriter, request *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	locations, err := app.db.GetAllLocations(ctx)
	if err != nil{
		jsonResp(writer, http.StatusInternalServerError,  map[string]any{"error":err.Error()})
		return
	}

	jsonResp(writer, http.StatusOK, map[string]any{"locations": locations})
}
