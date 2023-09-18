package main

// Importing necessary packages from the Go standard library.
import (
	"fmt"               // Package for formatted I/O.
	"net/http"          // Package for creating HTTP servers and clients.
	"net/http/httptest" // Package for HTTP testing utilities.
	"testing"           // Package for writing Go tests.
)

// Define a test function named TestHelloWorld.
func TestHelloWorld(t *testing.T) {
	// Create a test HTTP server using httptest.NewServer.
	// It will use the handleHelloWorld function to handle incoming HTTP requests.
	testServer := httptest.NewServer(http.HandlerFunc(handleHelloWorld))
	defer testServer.Close()
	// Close the test server when the test function is done.

	// Create an HTTP client for sending requests to the test server.
	testClient := testServer.Client()

	// Print the URL of the test server to the console.
	fmt.Println(testServer.URL)

	// Send a GET request to the test server's URL using the test client.
	testClient.Get(testServer.URL)
}
