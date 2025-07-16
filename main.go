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
	http.HandleFunc("/refresh", middleware.CorsMiddleware(authhandlers.RefreshTokens))
	http.HandleFunc("/register", middleware.CorsMiddleware(authhandlers.RegisterHandler))
	http.HandleFunc("/login", middleware.CorsMiddleware(authhandlers.LoginHandler))
	http.HandleFunc("/logout", middleware.CorsMiddleware(
		middleware.AuthMiddleware(authhandlers.LogOutHandler)))

	// file handlers
	http.HandleFunc("/upload", middleware.CorsMiddleware(
		middleware.AuthMiddleware(filehandlers.UploadHandler)))

	// Start the server on port 8080
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
