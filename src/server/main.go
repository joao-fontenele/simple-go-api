package main

import (
	"context"
	"encoding/json"
	"fmt"
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

type validationHandler struct {
	next http.Handler
}

// this avoids possible name clashes with other packages that may be using a key
// Name in the context
type validationContextKey string

func newValidationHandler(next http.Handler) http.Handler {
	return validationHandler{next}
}

func (h validationHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var request helloWorldRequest
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)
	if err != nil {
		log.Printf("Could not unmarshal the body to json")
		http.Error(rw, "Bad request", http.StatusBadRequest)
		return
	}

	c := context.WithValue(r.Context(), validationContextKey("name"), request.Name)
	r = r.WithContext(c)

	h.next.ServeHTTP(rw, r)
}

type helloWorldHandler struct{}

func newHelloWorldHandler() http.Handler {
	return helloWorldHandler{}
}

func (h helloWorldHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	name := r.Context().Value(validationContextKey("name")).(string)
	response := helloWorldResponse{Message: "Hello " + name}

	encoder := json.NewEncoder(rw)
	encoder.Encode(response)
}

func main() {
	port := 8080

	handler := newValidationHandler(newHelloWorldHandler())
	http.Handle("/hello", handler)

	// serve static files
	http.Handle(
		// the trailing `/` means this route will handle both `/images` and
		// `/images/somefile.ext`
		"/images/",
		// strip prefix was needed, or else the whole path would be forwarded to the
		// next handler
		http.StripPrefix("/images/", http.FileServer(http.Dir("./img"))),
	)

	log.Printf("Server listening on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
