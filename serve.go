package main

import (
	"net/http"
)

func main() {
	// Serve static files from the "web" directory at the root URL.
	// This will make index.html and other assets available at http://localhost:8081/
	http.Handle("/", http.FileServer(http.Dir("web")))

	// Start the server on port 8081.
	// This keeps it separate from your backend (which runs on 8080).
	println("Serving frontend on http://localhost:8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err) // Print error and exit if the server fails to start.
	}
}

// To run the server, use the following command:
// go run serve.go
