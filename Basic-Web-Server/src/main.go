package main

import (
	"fmt"
	"log"
	"net/http"

	"Basic-Server/src/handlers"
)

func main() {
	// Serve static files from the "static" directory
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)                   // Serve the homepage and static files
	http.HandleFunc("/hello", handlers.HelloHandler) // Use the HelloHandler
	http.HandleFunc("/form", handlers.HandleForm)    // Use the HandleForm

	url := "http://localhost:8080"
	fmt.Println("Server started at port 8080: ", url)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
