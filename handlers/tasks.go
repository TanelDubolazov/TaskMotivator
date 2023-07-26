package handlers

import (
	"encoding/json"
	"net/http"

	"TaskMotivator/models"
)

// TasksHandler handles the request for tasks-related endpoints ("/tasks").
func TasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Implement the logic to fetch tasks from the database and return them as JSON
		tasks, err := models.GetTasks("")
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
