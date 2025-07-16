package authhandlers

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

		var Err db.Message

		if err != nil {
			var errorMessage db.Message
			errorMessage.Message = err.Error()
			w.WriteHeader(http.StatusNotAcceptable)
			Err.Message = "Inappropriate input"
			json.NewEncoder(w).Encode(Err)
		}
		// fmt.Printf("Data retrieved: %+v\n", user)

		// TEMPORARY BLOCK START
		for iter := 0; iter < len(db.Array); iter++ {
			if user.Username == db.Array[iter].Username {
				Err.Message = "User With Such Username Already Exists"
				w.WriteHeader(http.StatusForbidden) // 403
				json.NewEncoder(w).Encode(Err)
				return
			} else if db.Array[iter].Username == "" { // this is stupid, but temporary
				db.Array[iter] = user
				db.ArrayJWT[iter].Access = user.Username
				db.ArrayJWT[iter].Refresh = user.Password

				w.WriteHeader(http.StatusCreated)            // 201
				json.NewEncoder(w).Encode(db.ArrayJWT[iter]) // returns tokens. TODO: implement properly
				return
			}
		}
		// TEMPORARY BLOCK END
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
