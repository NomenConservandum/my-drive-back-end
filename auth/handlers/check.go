package handlers

import (
	"encoding/json"
	"net/http"
)

// Checks for the correct API key. Will be modified to do more, e.g JWT fetching
func CheckHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet: // temporary solution
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("")
	case http.MethodPost: // TODO: JWT fetching. To be implemented.
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("")
	}
}
