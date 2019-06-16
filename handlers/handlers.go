package handlers

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

type handler struct {
	client *redis.Client
}

/**
New creates a handler struct with an initialized client to the database
*/
func New(client *redis.Client) *handler {
	return &handler{client: client}
}

func Mocked() *handler {
	return &handler{}
}

func (h *handler) HelloWorld(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello World!\n"))
}

func (h *handler) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	_, _ = io.WriteString(w, `{"alive": true}`)
}

// Handler responsible for aggregating information per POST
func (h *handler) NodeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "node choosen: %v\n", vars["nodename"])
}

// Handler responsible for aggregating information per POST
func (h *handler) NodesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "node choosen: %v\n", vars["nodename"])
	_, _ = fmt.Fprintf(w, "node choosen: %v\n", vars["processname"])
}

// Handler responsible for returning information per GET
func (h *handler) AnalyticsHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello World!\n"))
}
