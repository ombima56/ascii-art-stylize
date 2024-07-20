package ascii

import (
	"encoding/json"
	"errors"
	"html/template"
	"net/http"
	"os"
	"strings"
	"time"
)

func ErrorHandler(w http.ResponseWriter, errMsg string, statusCode int) {
	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		LogErrorToFile("Error parsing error template: " + err.Error())
		return
	}

	data := struct {
		StatusCode int
		ErrMsg     string
	}{
		StatusCode: statusCode,
		ErrMsg:     errMsg,
	}

	w.WriteHeader(statusCode)
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		LogErrorToFile("Error executing error template: " + err.Error())
	}
}

type LogEntry struct {
	Timestamp string `json:timestamp`
	Message   string `json:"message"`
}

func LogErrorToFile(errorMessage string) error {
	logFile, err := os.OpenFile("errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer logFile.Close()

	logEntry := LogEntry{
		Timestamp: time.Now().Format(time.RFC3339),
		Message:   errorMessage,
	}
	entryJSON, err := json.Marshal(logEntry)
	if err != nil {
		return err
	}

	_, err = logFile.Write(entryJSON)
	if err != nil {
		return err
	}
	_, err = logFile.WriteString("\n")
	return err
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, "Page Not Found", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		ErrorHandler(w, "Page Not Found", http.StatusNotFound)
		LogErrorToFile("Error parsing template: " + err.Error())
		return
	}

	err2 := tmpl.Execute(w, nil)
	if err2 != nil {
		ErrorHandler(w, "Internal Server Error", http.StatusInternalServerError)
		LogErrorToFile("Error executing template: " + err2.Error())
		return
	}
}

func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorHandler(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/ascii-art" {
		ErrorHandler(w, "Page Not Found", http.StatusNotFound)
		return
	}

	message := r.FormValue("message")
	bannerfile := r.FormValue("bannerfile")
	if message == "" || bannerfile == "" {
		ErrorHandler(w, "Bad Request: Missing message or banner file", http.StatusBadRequest)
		return
	}

	// Construct the file path
	filePath := "bannerfiles/" + bannerfile + ".txt"

	// Check if the banner file exists and has not been altered
	_, err := FileCheck(filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			ErrorHandler(w, "Internal Server Error: Banner file not found", http.StatusInternalServerError)
			LogErrorToFile("Banner file not found: " + filePath)
		} else if err.Error() == "the banner file has been altered" {
			ErrorHandler(w, "Internal Server Error: An unexpected error occurred. Please try again later.", http.StatusInternalServerError)
			LogErrorToFile("Banner file altered: " + filePath)
		} else {
			ErrorHandler(w, "Internal Server Error: An unexpected error occurred. Please try again later.", http.StatusInternalServerError)
			LogErrorToFile("Error with banner file: " + filePath)
		}
		return
	}

	data := strings.Split(message, "\r\n")
	var asciified strings.Builder
	for _, ch := range data {
		result, err := PrintBanner(ch, bannerfile)
		if err != nil {
			ErrorHandler(w, "Bad Request: Please use valid characters. Only printable characters from the ASCII table are allowed.", http.StatusBadRequest)
			LogErrorToFile("Error printing banner: " + err.Error())
			return
		}
		asciified.WriteString(result)
	}

	Data := struct {
		Ans   string
		Input string
	}{
		Ans:   asciified.String(),
		Input: message,
	}
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		ErrorHandler(w, "Internal Server Error: An unexpected error occurred. Please try again later.", http.StatusInternalServerError)
		LogErrorToFile("Error parsing template: " + err.Error())
		return
	}

	err = tmpl.Execute(w, Data)
	if err != nil {
		ErrorHandler(w, "Internal Server Error: An unexpected error occurred. Please try again later.", http.StatusInternalServerError)
		LogErrorToFile("Error executing template: " + err.Error())
	}
}
