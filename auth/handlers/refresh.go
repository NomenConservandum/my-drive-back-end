package authhandlers

import (
	"encoding/json"
	"myDrive/auth"
	"myDrive/db"
	"myDrive/utils"
	"net/http"
)

func RefreshTokens(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:

		defer r.Body.Close()
		tokens, err := utils.JsonToTokens(r)

		var Err db.Message

		if err != nil {
			var errorMessage db.Message
			errorMessage.Message = err.Error()
			w.WriteHeader(http.StatusNotAcceptable)
			Err.Message = "Inappropriate input"
			json.NewEncoder(w).Encode(Err)
		}

		// try to find the refresh token in the db
		// success: generate new access token
		// failure: send an error
		res, user := auth.IsValidRefresh(tokens.Refresh)
		if res {
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(auth.GenerateJWT(user))
		} else {
			w.WriteHeader(http.StatusForbidden)
			Err.Message = "ERROR"
			json.NewEncoder(w).Encode(Err)
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
