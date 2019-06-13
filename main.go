package main

import (
	"github.com/gorilla/mux"
	"github.com/nammn/node-aggregation/handlers"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", handlers.HelloWorld)
	r.HandleFunc("/health", handlers.HealthCheckHandler)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
