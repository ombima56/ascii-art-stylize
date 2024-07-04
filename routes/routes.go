package routes

import (
	"net/http"

	Ascii "ascii-art-stylize/ascii"
)

func RoutesSetUp(mux *http.ServeMux) {
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/", Ascii.IndexHandler)
	mux.HandleFunc("/submit", Ascii.SubmitHandler)
}
