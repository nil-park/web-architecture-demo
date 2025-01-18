package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/nil-park/web-architecture-demo/backend/pkg/k8s"
)

// versionHandler handles requests to the /version endpoint.
func versionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"version": "1.0.0",
	})
}

func main() {
	http.HandleFunc("/version", versionHandler)
	http.HandleFunc("/nodes", k8s.NodesHandler)
	http.HandleFunc("/pods", k8s.PodsHandler)

	log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
