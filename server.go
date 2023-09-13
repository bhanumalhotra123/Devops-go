// We are creating an API to understand its lifecycle.

// The "package main" statement indicates that this is the main package of our program.
package main

// We import necessary libraries or packages here.
// "log" for logging messages, and "net/http" for creating an HTTP server.
import (
	"fmt"
	"log"
	"net/http"
)

// The main function is the entry point of our program.
func main() {
	// We set up routes (endpoints) for our API.
	// These routes map to specific functions that handle HTTP requests.

	// This route "/hello-world" will be handled by the "handleHelloWorld" function.
	http.HandleFunc("/hello-world", handleHelloWorld)

	// This route "/health" will be handled by the "handleHealth" function.
	http.HandleFunc("/health", handleHealth)

	// We specify the address where our server will listen.
	addr := "localhost:8000" // This is the address and port we'll use during development.

	// We print a message to indicate that the server is listening.
	log.Printf("Listening on %s ...", addr)

	// We start the HTTP server and pass "nil" as the handler.
	err := http.ListenAndServe(addr, nil)

	// If there's an error starting the server, we log it and exit the program.
	if err != nil {
		log.Fatal(err)
	}
}

// The "handleHelloWorld" function handles requests to the "/hello-world" endpoint.
func handleHelloWorld(writer http.ResponseWriter, request *http.Request) {
	// This function is currently empty and doesn't do anything.
	// It's a placeholder for handling the "/hello-world" endpoint.
	if request.Method != "GET" {
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)

	}

	response := []byte("Hello World!!")
	_, err := writer.Write(response)
	if err != nil {
		fmt.Println(err)
	}

}

// The "handleHealth" function handles requests to the "/health" endpoint.
func handleHealth(writer http.ResponseWriter, request *http.Request) {
	// This function is also currently empty and doesn't do anything.
	// It's a placeholder for handling the "/health" endpoint.
	if request.Method != "GET" {
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)

	}

}
