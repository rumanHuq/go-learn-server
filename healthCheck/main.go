package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

// controller/multiplexer
func healthCheck(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()
	io.WriteString(w, currentTime.String())
}
func main() {
	// get route handler
	http.HandleFunc("/", healthCheck)
	// listener with error handler, 2nd param nil => using default multiPlexer
	log.Fatal(http.ListenAndServe(":8080", nil))
}
