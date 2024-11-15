package server

import (
	"encoding/json"
	"net/http"
)

func jsonResp(w http.ResponseWriter, status int, data map[string]any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
