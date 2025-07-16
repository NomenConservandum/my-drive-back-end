package handlers

import (
	"encoding/json"
	"myDrive/db"
	"myDrive/utils"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
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
		// fmt.Printf("Data retrieved: %+v\n", user)

		var Err db.Error

		// TEMPORARY BLOCK START
		for iter := 0; iter < len(db.Array); iter++ {
			if user.Username == db.Array[iter].Username {
				Err.Message = "User With Such Username Already Exists"
				w.WriteHeader(http.StatusForbidden) // 403
				json.NewEncoder(w).Encode(Err)
				return
			} else if db.Array[iter].Username == "" {
				db.Array[iter] = user
				w.WriteHeader(http.StatusCreated) // 201
				json.NewEncoder(w).Encode(user)   // returns a user, will return tokens later on. TODO: implement
				return
			}
		}
		// TEMPORARY BLOCK END
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
