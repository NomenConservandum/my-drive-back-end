package main

import (
	"fmt"
	authhandlers "myDrive/auth/handlers"
	"myDrive/auth/middleware"
	filehandlers "myDrive/files/handlers"
	"net/http"
)

func main() {
	// Define a handler function for the root path "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	// auth handlers
	http.HandleFunc("/check", middleware.CorsMiddleware(authhandlers.CheckHandler))
	http.HandleFunc("/register", middleware.CorsMiddleware(authhandlers.RegisterHandler))
	http.HandleFunc("/login", middleware.CorsMiddleware(authhandlers.LoginHandler))

	// file handlers
	http.HandleFunc("/upload", middleware.CorsMiddleware(filehandlers.UploadHandler))

	// Start the server on port 8080
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
