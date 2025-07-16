package authhandlers

import (
	"encoding/json"
	"net/http"
)

func LogOutHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(http.StatusOK)  // 200
		json.NewEncoder(w).Encode("") // returns nothing
		return
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
