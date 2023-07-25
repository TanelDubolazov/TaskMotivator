package main

import (
	"fmt"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Task Motivator!")
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	// I will implement the logic to handle tasks here
	// For example, fetch tasks from the database and return them as JSON
}

func main() {
	// Set up the routs and corresponding handlers
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/tasks", taskHandler)

	// Start the server on port 8080
	fmt.Println("Server started on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting the server : ", err)
	}
}
