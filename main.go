package main

import (
	"ascii-art-stylize/routes"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	port := "8080" // Default port
	if len(os.Args) > 2 {
		log.Fatal("Usage: go run main.go [port]")
		return
	}
	if len(os.Args) == 2 {
		port = os.Args[1]
		parsedPort, err := strconv.Atoi(port)
		if err != nil || parsedPort < 1024 || parsedPort > 65535 {
			log.Fatalf("Invalid port number: %v. It should be between 1024 and 65535.", port)
			return
		}
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	routes.RoutesSetUp(mux)
	log.Printf("Server listening on http://localhost:%s", port)
	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
