package handler

import (
	"net/http"
	"fmt"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/plain")

	fmt.Fprintf(w, "Hello, World!")
}
