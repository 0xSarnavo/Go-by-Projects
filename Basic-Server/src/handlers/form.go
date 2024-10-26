package handlers

import (
	"html/template"
	"log"
	"net/http"
)

// Citizen struct to hold form data
type Citizen struct {
	Name       string
	Species    string
	Planet     string
	Age        string
	Occupation string
	Starship   string
}

// HandleForm processes form submissions
func HandleForm(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		handleError(w, err, "Failed to parse form")
		return
	}

	citizen := Citizen{
		Name:       r.FormValue("name"),
		Species:    r.FormValue("species"),
		Planet:     r.FormValue("planet"),
		Age:        r.FormValue("age"),
		Occupation: r.FormValue("occupation"),
		Starship:   r.FormValue("starship"),
	}

	tmpl := template.Must(template.ParseFiles("static/output.html"))
	if err := tmpl.Execute(w, citizen); err != nil {
		handleError(w, err, "Failed to render output")
	}
}

// handleError logs the error and sends an HTTP error response
func handleError(w http.ResponseWriter, err error, message string) {
	log.Println(err) // Log the error
	http.Error(w, message, http.StatusInternalServerError) // Send error response
}
