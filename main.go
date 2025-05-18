package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from Go! 🌐 Host: %s\n", os.Getenv("HOSTNAME"))
	})

	fmt.Println("Server running on port", port)
	http.ListenAndServe(":"+port, nil)
}
