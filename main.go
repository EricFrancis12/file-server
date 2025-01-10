package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// Set the port on which the server will run
	port := "8080"

	// Get the current working directory (the folder the application is run in)
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	// Create a file server that will serve files from the current directory
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Get the file path relative to the current directory
		filePath := filepath.Join(dir, r.URL.Path)

		// Check if the requested file exists
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			w.Write([]byte("Not found"))
			return
		}

		// Serve the file at the requested path
		http.ServeFile(w, r, filePath)
	})

	// Print the URL the server will be running on
	fmt.Printf("Serving files from %s at http://localhost:%s\n", dir, port)

	// Start the server
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}
