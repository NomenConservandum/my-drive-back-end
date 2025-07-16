package handlers

import (
	"encoding/json"
	"myDrive/db"
	"myDrive/utils"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		defer r.Body.Close()
		user, err := utils.JsonToUser(r)
		if err != nil {
			var errorMessage db.Error
			errorMessage.Message = err.Error()
			w.WriteHeader(http.StatusNotAcceptable)
			json.NewEncoder(w).Encode(user)
		}

		var Err db.Error

		// TEMPORARY BLOCK START
		for iter := 0; iter < len(db.Array); iter++ {
			if user == db.Array[iter] {
				w.WriteHeader(http.StatusCreated) // 201
				json.NewEncoder(w).Encode(user)   // returns a user, will return tokens later on. TODO: implement
				return
			} else if user.Username == db.Array[iter].Username {
				Err.Message = "Password Does Not Match"
				w.WriteHeader(http.StatusForbidden) // 403
				json.NewEncoder(w).Encode(Err)
				return
			}
		}
		// TEMPORARY BLOCK END

		Err.Message = "User Not Found"

		w.WriteHeader(http.StatusNotFound) // 404
		json.NewEncoder(w).Encode(Err)
		return
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
