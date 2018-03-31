// proutils-api/handler_error
//
// API handlers for errors.

package main

import (
	"net/http"
)

// ErrorMethodNotAllowed renders a method not allowed response for invalid request
// types.
func ErrorMethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(405)
}

// ErrorNotFound renders a not found response for invalid API endpoints.
func ErrorNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
}
