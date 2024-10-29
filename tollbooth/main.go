package main

import (
	"encoding/json"
	"log"
	"net/http"

	tollbooth "github.com/didip/tollbooth/v7"
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
	message := Message{
		Status: "Request Failed",
		Body:   "API is at capacity",
	}
	jsonMessage, _ := json.Marshal(message)
	tlbthLimiter := tollbooth.NewLimiter(1, nil)
	tlbthLimiter.SetMessageContentType("application/json")
	tlbthLimiter.SetMessage(string(jsonMessage))

	http.Handle("/ping", tollbooth.LimitHandler(tlbthLimiter, http.HandlerFunc(endpointHandler)))

	log.Println("Server is running on port localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("Error starting server", err)
	}
}
