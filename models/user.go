package models

import (
	"database/sql"
	"time"
)

// User represents a single user in the application
type User struct {
	ID        int       // Unique identifier for the user
	FirstName string    // First name of the user
	LastName  string    // Last name of the user
	Email     string    // Email address of the user
	CreatedAt time.Time // Timestamp of when the user was created
}

// NewUser adds a new user to the database
func NewUser(user *User) error {
	// Open the database connection
	db, err := sql.Open("sqlite3", "taskmotivator.db")
	if err != nil {
		return err
	}
	defer db.Close()

	// Prepare the SQL Insert query
	insertQuery := "INSERT INTO users (first_name, last_name, email, created_at) VALUES (?,?,?,?)"

	// Execute the query with the user details
	_, err = db.Exec(insertQuery, user.FirstName, user.LastName, user.Email, user.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

// GetUserByID fetches a user from the database by their ID
func GetUserByID(userID int) (*User, error) {
	// Open the database connection
	db, err := sql.Open("sqlite3", "taskmotivator.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Prepare the SQL SELECT query
	selectQuery := "SELECT id, first_name, last_name, email, created_at FROM users WHERE id = ?"

	// Execute the query and get the result
	row := db.QueryRow(selectQuery, userID)

	// Create a new User instance to hold the result
	var user User

	// Scan the result into the user instance
	err = row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateUser updates an existing user in the database
func UpdateUser(user *User) error {
	// Open the database connection
	db, err := sql.Open("sqlite3", "taskmotivator.db")
	if err != nil {
		return err
	}
	defer db.Close()

	// Prepare the SQL UPDATE query
	updateQuery := "UPDATE users SET first_name=?, last_name=?, email=?, created_at=? WHERE id=?"

	// Execute the query with the updated user details
	_, err = db.Exec(updateQuery, user.FirstName, user.LastName, user.Email, user.CreatedAt, user.ID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteUser deletes an existing user from the database
func DeleteUser(user *User) error {
	// Open the database connection
	db, err := sql.Open("sqlite3", "taskmotivator.db")
	if err != nil {
		return err
	}
	defer db.Close()

	// Prepare the SQL DELETE query
	deleteQuery := "DELETE FROM users WHERE id=?"

	// Execute the query to delete the user by their ID
	_, err = db.Exec(deleteQuery, user.ID)
	if err != nil {
		return err
	}

	return nil
}

// GetUsers fetches users from the database based on the provided criteria
func GetUsers(criteria string) ([]*User, error) {
	// Open the database connection
	db, err := sql.Open("sqlite3", "taskmotivator.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Prepare the SQL SELECT query based on criteria (if provided)
	var selectQuery string
	if criteria != "" {
		selectQuery = "SELECT id, first_name, last_name, email, created_at FROM users WHERE " + criteria
	} else {
		selectQuery = "SELECT id, first_name, last_name, email, created_at FROM users"
	}

	// Execute the query and get the result
	rows, err := db.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Process the result and populate the users slice
	var users []*User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	// Check for any errors that occurred during the iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
