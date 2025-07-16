package main

import (
	"fmt"
	"myDrive/auth/handlers"
	"myDrive/auth/middleware"
	"net/http"
)

func main() {
	// Define a handler function for the root path "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	// auth handlers
	http.HandleFunc("/check", middleware.CorsMiddleware(handlers.CheckHandler))
	http.HandleFunc("/register", middleware.CorsMiddleware(handlers.RegisterHandler))
	http.HandleFunc("/login", middleware.CorsMiddleware(handlers.LoginHandler))

	// Start the server on port 8080
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
