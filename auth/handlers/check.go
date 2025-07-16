package authhandlers

import (
	"encoding/json"
	"myDrive/auth"
	"net/http"
)

// Checks for the correct API key. Will be modified to do more, e.g JWT fetching
func CheckHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet: // temporary solution
		// checking the access token
		if len(r.Header["Authorization"]) != 0 && auth.IsValidAccess(r.Header["Authorization"][0]) {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("")
		} else { // user needs to refresh the token
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode("")
		}
	}
}
