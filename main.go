package main

import (
	"github.com/gorilla/mux"
	"github.com/nammn/node-aggregation/database"
	"github.com/nammn/node-aggregation/handlers"
	"log"
	"net/http"
)

/**
Main class, responsible for creating a connection client and open the http server
*/
func main() {
	client, err := database.NewClient("")
	if err != nil {
		log.Fatal("problem connecting to the database")
	}
	pathHandlers := handlers.New(client)
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", pathHandlers.HelloWorld)
	r.HandleFunc("/health", pathHandlers.HealthCheckHandler)
	r.HandleFunc("/v1/metrics/node/{nodename}", pathHandlers.NodeHandler).Methods("POST")
	r.HandleFunc("/v1/metrics/nodes{nodename}/process/{processname}", pathHandlers.NodeHandler).Methods("POST")
	r.HandleFunc("/v1/analytics", pathHandlers.AnalyticsHandler).Methods("GET")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
