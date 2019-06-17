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
	Client database.RedisConnection
}

/**
NewClient creates a handler struct with an initialized Client to the database
*/
func NewHandler(client database.RedisConnection) *handler {
	return &handler{Client: client}
}

func (h *handler) HelloWorld(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello World!\n"))
}

func (h *handler) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err := h.Client.Ping()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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
	nodeStat.Timestamp = time.Now().Unix()
	err = h.Client.SaveNodeStatValue(nodeStat)
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
func (h *handler) AnalyticsNodesHandler(w http.ResponseWriter, r *http.Request) {

	_, _ = w.Write([]byte("Hello World!\n"))
}

// Handler responsible for returning information per GET
// Handlers: most recent POST
func (h *handler) AnalyticProcessesHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello World!\n"))
}

// Handler responsible for returning information per GET
// Handlers: most recent POST
func (h *handler) AnalyticSpecificProcessHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello World!\n"))
}
