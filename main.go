package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

var (
	// Redis client with connection pooling
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		PoolSize: 100, // Limit the pool size to avoid connection exhaustion
	})
	redisMutex = sync.Mutex{}
)

// Function to implement rate limiting with context for timeout
func isRateLimited(ip string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond) // Timeout to free up resources
	defer cancel()

	redisMutex.Lock()
	defer redisMutex.Unlock()

	count, err := redisClient.Get(ctx, ip).Int()
	if err != nil && err != redis.Nil {
		log.Printf("Redis error: %v", err)
		return false
	}

	if count >= 100 { // Example: Limit to 100 requests per minute
		return true
	}

	redisClient.Incr(ctx, ip)
	redisClient.Expire(ctx, ip, time.Minute)
	return false
}

// Optimized catch-all handler for non-existent endpoints
func dynamicEndpointHandler(w http.ResponseWriter, r *http.Request) {
	// Set a context with timeout to ensure handlers don’t hang
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	// Get the client's IP
	ip := r.RemoteAddr
	if isRateLimited(ip) {
		// Redirect to an amusing website if rate limit is exceeded
		http.Redirect(w, r, "https://www.pointerpointer.com", http.StatusSeeOther)
		return
	}

	// Simulate delay for realism, but respect the context
	select {
	case <-time.After(1 * time.Second):
		// Continue processing
	case <-ctx.Done():
		http.Error(w, "Request timed out", http.StatusGatewayTimeout)
		return
	}

	// Add random trolley headers
	addTrolleyHeaders(w)

	// Randomize status code
	statusCode := getRandomStatusCode()

	// Check if the request is attempting a common pattern (e.g., /api/v1/something)
	if strings.HasPrefix(r.URL.Path, "/api/v1/") {
		response := map[string]interface{}{
			"message":  getRandomComicalResponse(),
			"status":   "success",
			"endpoint": r.URL.Path,
			"data":     generateFakeData(r.URL.Path),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(response)
	} else {
		// If the request doesn't match the API pattern, return a 404
		http.NotFound(w, r)
	}
}

// Function to generate fake data based on the requested endpoint
func generateFakeData(endpoint string) map[string]string {
	return map[string]string{
		"info": "This is fake data generated for " + endpoint,
		"hint": "Nope, nothing valuable here. Move along!",
	}
}

func main() {
	// Use Gorilla Mux for routing
	r := mux.NewRouter()
	r.PathPrefix("/").HandlerFunc(dynamicEndpointHandler)

	// Create and configure the HTTP server
	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 5 * time.Second,  // Limit write duration
		ReadTimeout:  5 * time.Second,  // Limit read duration
		IdleTimeout:  15 * time.Second, // Keep-alive timeout
	}

	log.Println("H4X0R_B41T running on port 8080 with optimizations for stability...")
	log.Fatal(srv.ListenAndServe())
}
