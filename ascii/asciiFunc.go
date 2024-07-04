package Ascii

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
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

	err2 := tmpl.Execute(w, nil)
	if err2 != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err2)
		return
	}
}

func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/submit" {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		return
	}

	message := r.FormValue("message")
	bannerfile := r.FormValue("bannerfile")
	if message == "" || bannerfile == "" {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	data := strings.Split(message, "\r\n")
	var asciified string
	for _, ch := range data {
		asciified += PrintBanner(ch, bannerfile)
	}

	if asciified == "" {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	Data := struct {
		Ans string
	}{
		Ans: asciified,
	}
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		log.Printf("Error parsing template: %v", err)
		return
	}
	err2 := tmpl.Execute(w, Data)
	if err2 != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err2)
		return
	}
}
