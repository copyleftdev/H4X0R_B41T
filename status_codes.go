package main

import (
	"math/rand"
	"net/http"
	"time"
)

// List of possible status codes
var statusCodes = []int{
	http.StatusOK,                  // 200
	http.StatusCreated,             // 201
	http.StatusAccepted,            // 202
	http.StatusNoContent,           // 204
	http.StatusMovedPermanently,    // 301
	http.StatusBadRequest,          // 400
	http.StatusUnauthorized,        // 401
	http.StatusForbidden,           // 403
	http.StatusNotFound,            // 404
	http.StatusTeapot,              // 418
	http.StatusInternalServerError, // 500
	http.StatusNotImplemented,      // 501
	http.StatusServiceUnavailable,  // 503
}

// Function to get a random status code
func getRandomStatusCode() int {
	rand.Seed(time.Now().UnixNano())
	return statusCodes[rand.Intn(len(statusCodes))]
}
