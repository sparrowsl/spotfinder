package server

import (
	"encoding/json"
	"net/http"
)

func jsonResp(w http.ResponseWriter, status int, data map[string]any) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
