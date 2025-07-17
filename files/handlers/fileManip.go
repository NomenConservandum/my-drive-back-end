package handlers

import (
	"encoding/json"
	"fmt"
	"myDrive/db"
	"myDrive/utils"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func FileManipulationHandler(w http.ResponseWriter, r *http.Request) {
	// get the user
	username := utils.GetUserFromRequest(r).Username
	id, err := strconv.Atoi(strings.Split(r.URL.Path, "/")[2])
	if err != nil {
		w.WriteHeader(http.StatusConflict) // 409
		json.NewEncoder(w).Encode("Could not parse URL")
		return
	}
	if db.ArrayFiles[id].Owner != username { // user is not owner
		w.WriteHeader(http.StatusUnauthorized) // 401
		json.NewEncoder(w).Encode("")
		return
	}
	switch r.Method {
	case http.MethodGet:
		// Open the file
		filename := db.ArrayFiles[id].Name
		file, err := os.Open(filepath.Join("./uploads/", username, filename))
		if err != nil {
			if os.IsNotExist(err) {
				w.WriteHeader(http.StatusNotFound) // 404
				json.NewEncoder(w).Encode("No such file")
			} else {
				w.WriteHeader(http.StatusInternalServerError) // 500
				json.NewEncoder(w).Encode("")
			}
			return
		}
		defer file.Close()

		// Get file info for Content-Length
		fileInfo, err := file.Stat()
		if err != nil {
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		// Set headers
		w.Header().Set("Content-Disposition", "attachment; filename="+filename)
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Length", fmt.Sprint(fileInfo.Size()))

		// Stream the file to the client
		http.ServeContent(w, r, filename, fileInfo.ModTime(), file)
		return
	case http.MethodDelete:
		return
	}
}
