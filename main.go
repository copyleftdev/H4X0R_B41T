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
	"go.uber.org/zap"
)

// Redis client with connection pooling
var (
	redisClient *redis.Client
	redisMutex  sync.Mutex
	logger      *zap.Logger
)

// Color codes for terminal output
const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Cyan    = "\033[36m"
	Magenta = "\033[35m"
)

// Function to add color to log messages
func coloredLogger(level string, message string) string {
	switch level {
	case "INFO":
		return Green + message + Reset
	case "DEBUG":
		return Cyan + message + Reset
	case "WARN":
		return Yellow + message + Reset
	case "ERROR":
		return Red + message + Reset
	case "ALERT":
		return Magenta + message + Reset
	default:
		return message
	}
}

// Initialize Zap Logger
func initLogger() {
	var err error
	// Use the production configuration, which provides good defaults for performance
	logger, err = zap.NewProduction()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Sync() // Flushes buffer, if any
}

func initRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "redis:6379", // Assuming Redis runs on 'redis' in Docker
		PoolSize: 100,          // Limit the pool size to avoid connection exhaustion
	})
}

func main() {
	initLogger()
	initRedis()

	r := mux.NewRouter()
	r.PathPrefix("/").HandlerFunc(dynamicEndpointHandler)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 5 * time.Second,  // Limit write duration
		ReadTimeout:  5 * time.Second,  // Limit read duration
		IdleTimeout:  15 * time.Second, // Keep-alive timeout
	}

	logger.Info(coloredLogger("INFO", "🏁 H4X0R_B41T running on port 8080 with Zap logging enabled"))
	log.Fatal(srv.ListenAndServe())
}

func isRateLimited(ip string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	redisMutex.Lock()
	defer redisMutex.Unlock()

	count, err := redisClient.Get(ctx, ip).Int()
	if err != nil && err != redis.Nil {
		logger.Error(coloredLogger("ERROR", "❌ Redis error"), zap.String("ip", ip), zap.Error(err))
		return false
	}

	if count >= 100 {
		logger.Warn(coloredLogger("WARN", "⚠️ Rate limit exceeded"), zap.String("ip", ip))
		return true
	}

	redisClient.Incr(ctx, ip)
	redisClient.Expire(ctx, ip, time.Minute)
	return false
}

func dynamicEndpointHandler(w http.ResponseWriter, r *http.Request) {
	ip := r.RemoteAddr
	if isRateLimited(ip) {
		http.Redirect(w, r, "https://www.pointerpointer.com", http.StatusSeeOther)
		return
	}

	time.Sleep(1 * time.Second)

	logger.Info(coloredLogger("INFO", "🌐 Received request"),
		zap.String("method", r.Method),
		zap.String("url", r.URL.String()),
		zap.String("ip", ip),
	)

	addTrolleyHeaders(w)

	statusCode := getRandomStatusCode()

	if strings.HasPrefix(r.URL.Path, "/api/v1/") {
		response := map[string]interface{}{
			"message":  getRandomComicalResponse(),
			"status":   "success",
			"endpoint": r.URL.Path,
			"data":     generateFakeData(r.URL.Path),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		logger.Info(coloredLogger("INFO", "✅ Responding with status code"),
			zap.Int("statusCode", statusCode),
			zap.String("endpoint", r.URL.Path),
		)
		json.NewEncoder(w).Encode(response)
	} else {
		http.NotFound(w, r)
		logger.Warn(coloredLogger("WARN", "🔍 Endpoint not found"), zap.String("url", r.URL.String()))
	}
}

func generateFakeData(endpoint string) map[string]string {
	return map[string]string{
		"info": "This is fake data generated for " + endpoint,
		"hint": "Nope, nothing valuable here. Move along!",
	}
}
