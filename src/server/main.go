package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloWorldResponse struct {
	Message string `json:"message"` // `json:"message" renames the field for marshalling`
}

func main() {
	port := 8080

	http.HandleFunc("/hello", helloWorldHandler)

	log.Printf("Server listening on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("GET /hello\n")

	response := helloWorldResponse{Message: "Hello World"}
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	encoder.Encode(&response)
}
