package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Message struct {
	Status string `json:"status"`
	Body   string `json:"body"`
}

func endpointHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	message := Message{
		Status: "success",
		Body:   "Successfully pinged",
	}
	err := json.NewEncoder(writer).Encode(message)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/ping", rateLimiter(endpointHandler).ServeHTTP)
	fmt.Println("Server is running on port localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("Error starting server", err)
	}
}
