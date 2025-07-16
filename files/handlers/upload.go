package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"myDrive/db"
	"myDrive/utils"
	"net/http"
	"os"
	"time"
)

const maxUploadSize = 100 << 20 // 100 MB
const uploadBasePath = "./uploads/"

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		// get the user
		username := utils.GetUserFromRequest(r).Username
		uploadPath := uploadBasePath + username

		var Err db.Message

		// Limit upload size
		r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
		if err := r.ParseMultipartForm(maxUploadSize); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			Err.Message = "Inappropriate input"
			json.NewEncoder(w).Encode(Err)
			return
		}

		// Get file from form data
		file, handler, err := r.FormFile("file")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			Err.Message = "Invalid file"
			json.NewEncoder(w).Encode(Err)
			return
		}
		defer file.Close()

		// Create upload directory if not exists
		if err := os.MkdirAll(uploadPath, os.ModePerm); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			Err.Message = "Server error"
			json.NewEncoder(w).Encode(Err)
			return
		}

		// Create destination file
		dstPath := fmt.Sprintf("%s/%s", uploadPath, handler.Filename)
		dst, err := os.Create(dstPath)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			Err.Message = "Server error"
			json.NewEncoder(w).Encode(Err)
			return
		}
		defer dst.Close()

		// Copy file to destination
		if _, err := io.Copy(dst, file); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			Err.Message = "Server error"
			json.NewEncoder(w).Encode(Err)
			return
		}

		// add file metadata to the db
		for iter := 0; iter < db.UsersNum*db.FilesNum; iter++ {
			if db.ArrayFiles[iter].Name == "" {
				db.ArrayFiles[iter].ID = iter
				db.ArrayFiles[iter].CreatedAt = time.Now()
				db.ArrayFiles[iter].Name = handler.Filename
				db.ArrayFiles[iter].Owner = username
				db.ArrayFiles[iter].Size = handler.Size
				break
			}
		}

		var text db.Message // TODO: remove and rename Err
		text.Message = "File " + handler.Filename + " uploaded successfully"
		w.WriteHeader(http.StatusCreated) // 201
		json.NewEncoder(w).Encode(text)
	}
}
