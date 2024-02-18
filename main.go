package main

import (
	"fmt"
	"net/http"
)



func main() {
	const port = ":3000"
	server := http.NewServeMux()

	server.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	server.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world! from POST"))
	})

	fmt.Printf("Server is running at %s\n", port)
	http.ListenAndServe(port, server)
}