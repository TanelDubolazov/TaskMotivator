package models

import (
	"database/sql"
	"time"
)

// Task represents a single task in the application

type Task struct {
	ID          int       // Unique identifier for the task
	Title       string    // Title of the task
	Description string    // Description of the task
	DueDate     time.Time // Due date of the task
	Priority    int       // Priority level of the task(1-3, high-low)
	Status      string    // Status of the task ("completed", "in progress", "not started")
}

// NewTask adds a new task to the database
func NewTask(task *Task) error {
	// open the database connection
	db, err := sql.Open("sqlite3", "taskmotivator.db")
	if err != nil {
		return err
	}
	defer db.Close()

	// Preparing the SQL Insert query
	insertQuery := "INSERT INTO tasks (title, description, due_date, priority, status) VALUES (?,?,?,?,?)"

	// Execute the query with the task details
	_, err = db.Exec(insertQuery, task.Title, task.Description, task.DueDate, task.Priority, task.Status)
	if err != nil {
		return err
	}
	return nil
}

// GetTaskByID fetches a task from the database by its ID
func GetTaskByID(taskID int) (*Task, error) {
	// Open the database connection
	db, err := sql.Open("sqlite3", "taskmotivator.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Prepare the SQL SELECT query
	selectQuery := "SELECT id, title, description, due_date, priority, status FROM tasks WHERE id = ?"

	// Execute the query and get the result
	row := db.QueryRow(selectQuery, taskID)

	// Create a new Task instance to hold the result
	var task Task

	// Scan the result into the task instance
	err = row.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Priority, &task.Status)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

// UpdateTask updates an existing task in the database
func UpdateTask(task *Task) error {
	// Open the database connection
	db, err := sql.Open("sqlite3", "taskmotivator.db")
	if err != nil {
		return err
	}
	defer db.Close()

	// Prepare the SQL UPDATE query
	updateQuery := "UPDATE tasks SET title=?, description=?, due_date=?, priority=?, status=? WHERE id=?"

	// Execute the query with the updated task details
	_, err = db.Exec(updateQuery, task.Title, task.Description, task.DueDate, task.Priority, task.Status, task.ID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteTask deletes an existing task in the database
func DeleteTask(task *Task) error {
	// Open the database connection
	db, err := sql.Open("sqlite3", "taskmotivator.db")
	if err != nil {
		return err
	}
	defer db.Close()

	// Prepare the SQL DELETE query
	deleteQuery := "DELETE FROM tasks WHERE id=?"

	// Execute the query to delete the task by its ID
	_, err = db.Exec(deleteQuery, task.ID)
	if err != nil {
		return err
	}

	return nil
}

// GetTasks fetches the tasks from the database based on the provided criteria
func GetTasks(criteria string) ([]*Task, error) {
	// Open the database connection
	db, err := sql.Open("sqlite3", "taskmotivator.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Prepare the SQL SELECT query based on criteria (if provided)
	var selectQuery string
	if criteria != "" {
		selectQuery = "SELECT id, title, description, due_date, priority, status FROM tasks WHERE " + criteria
	} else {
		selectQuery = "SELECT id, title, description, due_date, priority, status FROM tasks"
	}

	// Execute the query and get the result
	rows, err := db.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Process the result and populate the tasks slice
	var tasks []*Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Priority, &task.Status)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	// Check for any errors that occurred during the iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
