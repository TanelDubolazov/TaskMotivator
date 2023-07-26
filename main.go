package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"TaskMotivator/models" // Import the models package
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Task Motivator!")
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Implement the logic to fetch tasks from the database and return them as JSON
		tasks, err := getTasksFromDB()
		if err != nil {
			http.Error(w, "Failed to fetch tasks from the database", http.StatusInternalServerError)
			return
		}

		// Convert tasks slice to JSON and write it to the response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getTasksFromDB() ([]models.Task, error) {
	// Open the database connection
	db, err := sql.Open("sqlite3", "taskmotivator.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Query the tasks from the database
	rows, err := db.Query("SELECT id, title, description, due_date, priority, status FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Process the result and populate the tasks slice
	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Priority, &task.Status)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	// Check for any errors that occurred during the iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func main() {
	// Set up the routes and corresponding handlers
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/tasks", taskHandler)

	// Start the server on port 8080
	fmt.Println("Server started on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting the server: ", err)
	}
}
