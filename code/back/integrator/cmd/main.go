// main.go

package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	gokithttp "github.com/go-kit/kit/transport/http"
	"log"
)

// Define the request and response structures
type helloRequest struct{}

type helloResponse struct {
	Message string `json:"message"`
}

// Create the endpoint
func makeHelloEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return helloResponse{Message: "Hello World from Go Kit!"}, nil
	}
}

// Decode the request (no parameters needed)
func decodeHelloRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	return helloRequest{}, nil
}

// Encode the response as JSON
func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func main() {
	// Create a handler for the endpoint
	helloHandler := gokithttp.NewServer(
		makeHelloEndpoint(),
		decodeHelloRequest,
		encodeResponse,
	)

	// Start the HTTP server
	http.Handle("/hello", helloHandler)
	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
