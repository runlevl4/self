package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const content = "Content-Type"
const contentJSON = "application/json"

//Health responds if the service is ready to take traffic.
func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(content, contentJSON)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"status": "HEALTHY"}`)
}

// InfoResponse provides the response details for the Info route.
type InfoResponse struct {
	Status    string `json:"status,omitempty"`
	Pod       string `json:"pod,omitempty"`
	PodIP     string `json:"pod_ip,omitempty"`
	Node      string `json:"node,omitempty"`
	Namespace string `json:"namespace,omitempty"`
}

//Info responds if the service is up.
func Info(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(content, contentJSON)
	w.WriteHeader(http.StatusOK)

	// Build response
	host, _ := os.Hostname()
	ir := InfoResponse{
		Status:    "UP",
		Pod:       host,
		PodIP:     os.Getenv("KUBERNETES_NAMESPACE_POD_IP"),
		Node:      os.Getenv("KUBERNETES_NODENAME"),
		Namespace: os.Getenv("KUBERNETES_NAMESPACE"),
	}
	rs, err := json.Marshal(ir)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "unable to parse info"}`)
	}
	fmt.Fprintf(w, string(rs))
}

// Hello is our simple entrypoint.
func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(content, contentJSON)
	w.WriteHeader(http.StatusOK)

	s := `{"message": "Hello, SELF!"}`

	fmt.Fprintf(w, s)
	fmt.Fprintf(w, "\n")
}
