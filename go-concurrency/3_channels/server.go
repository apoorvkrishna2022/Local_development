package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Response struct {
	Index string `json:"index"`
}

// processHandler simulates processing time between 2 to 10 seconds
func processHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("server is hit with index ", r.URL.Query().Get("index"))
	// Simulate processing time between 2 to 10 seconds
	rand.Seed(time.Now().UnixNano())
	processTime := rand.Intn(9) + 2
	time.Sleep(time.Duration(processTime) * time.Second)

	// Extract index from query parameter
	index := r.URL.Query().Get("index")

	// Create response object
	resp := Response{
		Index: index,
	}

	// Marshal response object to JSON
	responseJSON, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")

	// Write response JSON
	w.Write(responseJSON)
}

func main() {
	// Register handler for "/process" endpoint
	http.HandleFunc("/process", processHandler)

	// Start the server
	fmt.Println("Server started at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
