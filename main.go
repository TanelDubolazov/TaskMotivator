package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"TaskMotivator/models" // Import the models package

	_ "github.com/mattn/go-sqlite3"
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
	// Create the SQLite database and the tasks table if they don't exist
	db, err := sql.Open("sqlite3", "taskmotivator.db")
	if err != nil {
		log.Fatal("Error opening the database: ", err)
	}
	defer db.Close()

	createTableQuery := `
		CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			description TEXT,
			due_date DATETIME,
			priority INTEGER,
			status TEXT
		)
	`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal("Error creating the tasks table: ", err)
	}

	// Populate the database with some initial tasks
	task1 := models.Task{
		Title:       "Task 1",
		Description: "This is the first task",
		DueDate:     time.Now().AddDate(0, 0, 7), // Due date is 7 days from now
		Priority:    1,
		Status:      "not started",
	}
	err = models.NewTask(&task1)
	if err != nil {
		log.Fatal("Error adding task:", err)
	}

	task2 := models.Task{
		Title:       "Task 2",
		Description: "This is the second task",
		DueDate:     time.Now().AddDate(0, 0, 14), // Due date is 14 days from now
		Priority:    2,
		Status:      "in progress",
	}
	err = models.NewTask(&task2)
	if err != nil {
		log.Fatal("Error adding task:", err)
	}

	// Serve static files from the "static" directory
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Set up the routes and corresponding handlers
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/tasks", taskHandler)

	// Start the server on port 8080
	fmt.Println("Server started on port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting the server: ", err)
	}
}
