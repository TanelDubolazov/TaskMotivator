// handlers/home.go
package handlers

import (
	"net/http"
)

// HomeHandler handles the request to the home page ("/").
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// You can handle the request here and generate the response.
	// For example, you might render an HTML template or return a simple message.
	w.Write([]byte("Welcome to TaskMotivator!"))
}
