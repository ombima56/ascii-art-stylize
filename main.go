package main

import (
	"log"
	"net/http"

	"ascii-art-stylize/routes"
)

func main() {
	mux := http.NewServeMux()

	routes.RoutesSetUp(mux)
	log.Println("Server Listening on http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
