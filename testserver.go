package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Quote struct {
	Bid string `json:"bid"`
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	time.Sleep(100 * time.Millisecond) // Simulate delay
	response := Quote{Bid: "5.42"}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/cotacao", handleRequest)
	serverAddr := ":8080"
	log.Printf("Starting test server on %s", serverAddr)
	if err := http.ListenAndServe(serverAddr, nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
