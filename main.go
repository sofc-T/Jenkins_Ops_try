package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/whatsup", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "What's up world!")
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "OK")
	})

	fmt.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}