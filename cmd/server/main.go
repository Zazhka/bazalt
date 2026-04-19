package main

import (
	"log"
	"net/http"

	"github.com/Zazhka/bazalt/internal/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.HelloHandler)

	port := ":8080"
	log.Printf("Server starting on http://localhost%s", port)

	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
