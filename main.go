package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io"
	//"strings"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func enableCORS(w http.ResponseWriter) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

func main() {
	// Define a handler function for the root path "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		enableCORS(w)
		switch r.Method {
		case http.MethodPost:
			defer r.Body.Close()

			var user User

			bodyBytes, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Error reading request body", http.StatusBadRequest)
				return
			}
			
			// Convert bytes to string
			bodyString := string(bodyBytes)
			fmt.Println("Parsed JSON: " + bodyString)

			err = json.Unmarshal([]byte(bodyString), &user)
			if err != nil {
				fmt.Printf("Error parsing JSON: %v\n", err)
				return
			}
			fmt.Printf("Data retrieved: %+v\n", user)

			// Set the header
			w.WriteHeader(http.StatusCreated) // For POST requests
			// Encode the data to JSON and write to response
			json.NewEncoder(w).Encode(user)
			
    		//fmt.Fprint(w, "Success")
			return
		default:
			fmt.Println("Got an unappropriate request...\nThe method is: " + r.Method)
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Start the server on port 8080
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}