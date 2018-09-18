package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type helloWorldResponse struct {
	Message string `json:"message"` // `json:"message" renames the field for marshalling`
}

// defining this type is optional for unmarshalling, but for clarity of the code
// it's better to define it. It also easier than manually cast fields
type helloWorldRequest struct {
	Name string `json:"name"`
}

func main() {
	port := 8080

	http.HandleFunc("/hello", helloWorldHandler)

	log.Printf("Server listening on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v /hello\n", r.Method)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body stream")
		http.Error(w, "Bad request", http.StatusBadRequest)
		// Body.Close(), to close the request stream, was not called because
		// apparently ServeHTTP handler already does that automatically
		return
	}

	var request helloWorldRequest
	err = json.Unmarshal(body, &request)
	if err != nil {
		log.Printf("Could not unmarshal the body to json")
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	response := helloWorldResponse{Message: "Hello " + request.Name}
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	encoder.Encode(&response)
}
