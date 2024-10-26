package handlers

import (
	"fmt"
	"net/http"
)

// HelloHandler handles the /hello endpoint
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Use GET Method", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintln(w, "Hello Citizen")
}
