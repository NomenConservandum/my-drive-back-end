package handlers

import (
	"encoding/json"
	"myDrive/db"
	"myDrive/utils"
	"net/http"
)

func GetFilesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// get the user
		username := utils.GetUserFromRequest(r).Username

		var files []db.Metadata

		// add file metadata to the db
		for iter := 0; iter < db.UsersNum*db.FilesNum; iter++ {
			if db.ArrayFiles[iter].Owner == username {
				//println(db.ArrayFiles[iter].Name + " is owned by " + db.ArrayFiles[iter].Owner)
				files = append(files, db.ArrayFiles[iter])
			}
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(files)
	}
}
