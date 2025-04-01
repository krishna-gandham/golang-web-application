package main

import (
	"fmt"
	"net/http"
	"time"
)

// Handler function to return the current date and time
func handler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Fprintf(w, "Current Date & Time: %s", currentTime)
}

func main() {
	http.HandleFunc("/", handler) // Route the root URL to the handler
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil) // Start the web server on port 8080
}
