package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: .env file not found or couldn't be loaded.")
	}

	val := os.Getenv("WHATSUP_VALUE")
	if val == "" {
		//  if no whatsup value is set, use boring old 
		val = "Hello, World!"
	}

	http.HandleFunc("/whatsup", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, val)
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
