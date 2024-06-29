package main

import (
	Ascii "ascii-art-stylize/ascii"
	"html/template"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/submit", submitHandler)

	log.Println("Server Listening on http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		log.Printf("Error parsing template: %v", err)
		return
	}

	err2 := tmpl.Execute(w, r)
	if err2 != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err2)
		return
	}
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "500 Bad Request Method", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/submit" {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		return
	}

	message := r.FormValue("message")
	bannerfile := r.FormValue("bannerfile")

	data := Ascii.PrintBanner(message, bannerfile)
	if data == "" {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(data))
}
