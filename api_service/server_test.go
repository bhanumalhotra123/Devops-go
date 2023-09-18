package main

// Importing necessary packages from the Go standard library.
import (
	"fmt" // Package for formatted I/O.
	"io"
	"net/http"          // Package for creating HTTP servers and clients.
	"net/http/httptest" // Package for HTTP testing utilities.
	"strings"
	"testing" // Package for writing Go tests.

	"github.com/stretchr/testify/assert"
)

// Define a test function named TestHelloWorld.
func TestHelloWorldSuccess(t *testing.T) {
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
	response, err := testClient.Get(testServer.URL)
	if err != nil {
		t.Error(err)
	}

	// Read the response body.
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}
	// Close the response body after reading it.
	response.Body.Close()

	// Assert that the response body is "Hello World".
	assert.Equal(t, "Hello-World", string(bodyBytes))

	// Assert that the HTTP status code is 200 (OK).
	assert.Equal(t, 200, response.StatusCode)

}

func TestHelloWorldFailure(t *testing.T) {
	// Create a test HTTP server using httptest.NewServer.
	// It will use the handleHelloWorld function to handle incoming HTTP requests.
	testServer := httptest.NewServer(http.HandlerFunc(handleHelloWorld))
	defer testServer.Close()
	// Close the test server when the test function is done.

	// Create an HTTP client for sending requests to the test server.
	testClient := testServer.Client()

	// Print the URL of the test server to the console.
	fmt.Println(testServer.URL)

	body := strings.NewReader("some body")

	// Send a GET request to the test server's URL using the test client.
	response, err := testClient.Post(testServer.URL, "application/json", body)
	if err != nil {
		t.Error(err)
	}

	// Assert that the HTTP status code is 200 (OK).
	assert.Equal(t, 405, response.StatusCode)
}

func TestHealthSuccess(t *testing.T) {
	// Create a test HTTP server using httptest.NewServer.
	// It will use the handleHelloWorld function to handle incoming HTTP requests.
	testServer := httptest.NewServer(http.HandlerFunc(handleHealth))
	defer testServer.Close()
	// Close the test server when the test function is done.

	// Create an HTTP client for sending requests to the test server.
	testClient := testServer.Client()

	// Print the URL of the test server to the console.
	fmt.Println(testServer.URL)

	// Send a GET request to the test server's URL using the test client.
	response, err := testClient.Get(testServer.URL)
	if err != nil {
		t.Error(err)
	}

	// Read the response body.
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}
	// Close the response body after reading it.
	response.Body.Close()

	// Assert that the response body is "Hello World".
	assert.Equal(t, "OK", string(bodyBytes))

	// Assert that the HTTP status code is 200 (OK).
	assert.Equal(t, 200, response.StatusCode)

}

func TestHealthFail(t *testing.T) {
	// Create a test HTTP server using httptest.NewServer.
	// It will use the handleHelloWorld function to handle incoming HTTP requests.
	testServer := httptest.NewServer(http.HandlerFunc(handleHealth))
	// Close the test server when the test function is done.
	defer testServer.Close()
	// Close the test server when the test function is done.

	// Create an HTTP client for sending requests to the test server.
	testClient := testServer.Client()

	// Print the URL of the test server to the console.
	fmt.Println(testServer.URL)

	body := strings.NewReader("some body")

	// Send a GET request to the test server's URL using the test client.
	response, err := testClient.Post(testServer.URL, "application/json", body)
	if err != nil {
		t.Error(err)
	}

	// Assert that the HTTP status code is 200 (OK).
	assert.Equal(t, 405, response.StatusCode)
}
