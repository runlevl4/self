package handler

import (
	"fmt"
	"net/http"
)

const content = "Content-Type"
const contentJSON = "application/json"

//Health responds if the service is ready to take traffic.
func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(content, contentJSON)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"status": "HEALTHY"}`)
}

//Info responds if the service is up.
func Info(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(content, contentJSON)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"status": "UP"}`)
}

// Hello is our simple entrypoint.
func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(content, contentJSON)
	w.WriteHeader(http.StatusOK)

	s := `{"message": "Hello, SELF!"}`

	fmt.Fprintf(w, s)
	fmt.Fprintf(w, "\n")
}
