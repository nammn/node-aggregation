package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nammn/node-aggregation/database"
	"io"
	"net/http"
	"time"
)

type handler struct {
	client *database.RedisClient
}

/**
NewClient creates a handler struct with an initialized client to the database
*/
func New(client *database.RedisClient) *handler {
	return &handler{client: client}
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
// Saves Information per node as timeslice, cpu and mem
// Saved datastructure will be a struct
/**
■ “timeslice”: (float) number of seconds this measurement represents
■ “cpu”: (float) percentage used
■ “mem”: (float) percentage used
*/
func (h *handler) NodeHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var nodeStat database.NodeStat
	err := decoder.Decode(&nodeStat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	vars := mux.Vars(r)
	nodeStat.NodeName = vars["nodename"]
	nodeStat.Timestamp = time.Now()
	err = h.client.SaveNodeStatValue(nodeStat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "node choosen: %v\n", vars["nodename"])
}

// Handler responsible for aggregating information per POST
// Input is a JSON Body with timeslice,cpu,mem
func (h *handler) NodesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "node choosen: %v\n", vars["nodename"])
	_, _ = fmt.Fprintf(w, "node choosen: %v\n", vars["processname"])
}

// Handler responsible for returning information per GET
// Handlers: most recent POST
//
func (h *handler) AnalyticsHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello World!\n"))
}
