package models

import (
	"database/sql"
	"time"
)

// Reminder represents a single reminder in the application.
type Reminder struct {
	ID          int       // Unique ID for the reminder
	Title       string    // Title of the reminder
	Description string    // Description of the reminder
	DueDate     time.Time // Due date of the reminder
}

// NewReminder adds a new reminder to the database
func NewReminder(reminder *Reminder) error {
	// Open the database connection
	db, err := sql.Open("sqlite3", "taskmotivator.db")
	if err != nil {
		return err
	}
	defer db.Close()

	// Prepare the SQL Insert query
	insertQuery := "INSERT INTO reminders (title, description, due_date) VALUES (?,?,?)"

	// Execute the query with the reminder details
	_, err = db.Exec(insertQuery, reminder.Title, reminder.Description, reminder.DueDate)
	if err != nil {
		return err
	}

	return nil
}

// GetReminderByID fetches a reminder from the database by its ID
func GetReminderByID(reminderID int) (*Reminder, error) {
	// Open the database connection
	db, err := sql.Open("sqlite3", "taskmotivator.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Prepare the SQL SELECT query
	selectQuery := "SELECT id, title, description, due_date FROM reminders WHERE id = ?"

	// Execute the query and get the result
	row := db.QueryRow(selectQuery, reminderID)

	// Create a new Reminder instance to hold the result
	var reminder Reminder

	// Scan the result into the reminder instance
	err = row.Scan(&reminder.ID, &reminder.Title, &reminder.Description, &reminder.DueDate)
	if err != nil {
		return nil, err
	}

	return &reminder, nil
}

// UpdateReminder updates an existing reminder in the database
func UpdateReminder(reminder *Reminder) error {
	// Open the database connection
	db, err := sql.Open("sqlite3", "taskmotivator.db")
	if err != nil {
		return err
	}
	defer db.Close()

	// Prepare the SQL UPDATE query
	updateQuery := "UPDATE reminders SET title=?, description=?, due_date=? WHERE id=?"

	// Execute the query with the updated reminder details
	_, err = db.Exec(updateQuery, reminder.Title, reminder.Description, reminder.DueDate, reminder.ID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteReminder deletes an existing reminder from the database
func DeleteReminder(reminder *Reminder) error {
	// Open the database connection
	db, err := sql.Open("sqlite3", "taskmotivator.db")
	if err != nil {
		return err
	}
	defer db.Close()

	// Prepare the SQL DELETE query
	deleteQuery := "DELETE FROM reminders WHERE id=?"

	// Execute the query to delete the reminder by its ID
	_, err = db.Exec(deleteQuery, reminder.ID)
	if err != nil {
		return err
	}

	return nil
}

// GetReminders fetches reminders from the database based on the provided criteria
func GetReminders(criteria string) ([]*Reminder, error) {
	// Open the database connection
	db, err := sql.Open("sqlite3", "taskmotivator.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Prepare the SQL SELECT query based on criteria (if provided)
	var selectQuery string
	if criteria != "" {
		selectQuery = "SELECT id, title, description, due_date FROM reminders WHERE " + criteria
	} else {
		selectQuery = "SELECT id, title, description, due_date FROM reminders"
	}

	// Execute the query and get the result
	rows, err := db.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Process the result and populate the reminders slice
	var reminders []*Reminder
	for rows.Next() {
		var reminder Reminder
		err := rows.Scan(&reminder.ID, &reminder.Title, &reminder.Description, &reminder.DueDate)
		if err != nil {
			return nil, err
		}
		reminders = append(reminders, &reminder)
	}

	// Check for any errors that occurred during the iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return reminders, nil
}
