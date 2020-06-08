package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"	
	"github.com/runlevl4/self/api/internal/handler"
)

func main() {
	if err := run(); err != nil {
		log.Printf("error|shutting down|%s", err)
		os.Exit(1)
	}
}

func run() error {
	var log = log.New(os.Stdout, "self|", log.LstdFlags|log.Lmicroseconds)

	r := mux.NewRouter()
	r.HandleFunc("/hello", loggingMiddleware(handler.Hello)).Methods("GET")
	r.HandleFunc("/health", loggingMiddleware(handler.Health)).Methods("GET")
	r.HandleFunc("/info", loggingMiddleware(handler.Info)).Methods("GET")

	log.Printf("%s : Starting server on port 8000", time.Now().Format(time.RFC3339))
	log.Fatal(http.ListenAndServe(":8000", r))

	return nil
}

// HandleFunc provides our logging middleware functionality.
type HandleFunc func(w http.ResponseWriter, r *http.Request)

func loggingMiddleware(handler HandleFunc) HandleFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s : %s : %s", time.Now().Format(time.RFC3339), r.Method, r.URL)
		handler(w, r)
	}
	return fn
}
