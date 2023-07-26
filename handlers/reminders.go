package handlers

import (
	"encoding/json"
	"net/http"

	"Taskmotivator/models"
)

// RemindersHandler handles the request for reminders-related endpoints ("/reminders").
func RemindersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Implement the logic to fetch reminders from the database and return them as JSON
		reminders, err := models.GetReminders("")
		if err != nil {
			http.Error(w, "Failed to fetch reminders from the database", http.StatusInternalServerError)
			return
		}

		// Convert reminders slice to JSON and write it to the response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(reminders)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
