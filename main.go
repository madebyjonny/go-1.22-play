package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)


type Task struct {
    Task string `json:"task"`
}

var tasks []Task = []Task{}

func getenv(key, fallback string) string {
    value := os.Getenv(key)
    if len(value) == 0 {
        return fallback
    }
    return value
}

func main() {
	var port = ":" + getenv("PORT", "3000") 
	server := http.NewServeMux()

	server.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		jsonResponse, err := json.Marshal(tasks)
		if err != nil {
			http.Error(w, "Failed to marshal tasks to JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	})

	server.HandleFunc("POST /add/", func(w http.ResponseWriter, r *http.Request) {
		var requestBody Task
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&requestBody); err != nil {
			http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
			return
		}

		tasks = append(tasks, requestBody)

		jsonResponse, err := json.Marshal(tasks)
		if err != nil {
			http.Error(w, "Failed to marshal task to JSON", http.StatusInternalServerError)
			return
		}

		// Set content type header
		w.Header().Set("Content-Type", "application/json")

		// Write JSON response
		w.WriteHeader(http.StatusCreated)
		w.Write(jsonResponse)
	})

	fmt.Printf("Server is running at %s\n", port)
	log.Fatal(http.ListenAndServe(port, server))
}