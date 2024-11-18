package main

import (
	"math/rand"
	"net/http"
	"time"
)

// List of trolley headers
var trolleyHeaders = []struct {
	Key   string
	Value string
}{
	{"X-Powered-By", "Black Magic v1.0"},
	{"Server", "Spaghetti Code/2.0"},
	{"X-Robots-Tag", "please-hack-me"},
	{"X-Developer-Note", "This API was built during a 48-hour coding binge."},
	{"X-Humor", "Still better than your ex's cooking."},
	{"X-Obfuscation-Level", "Expert"},
	{"X-Debug-Token", "NULL_POINTER_EXCEPTION_12345"},
	{"X-Encryption", "ROT13 (just kidding, or are we?)"},
	{"X-Server-OS", "Windows 95"},
}

// Function to add random trolley headers
func addTrolleyHeaders(w http.ResponseWriter) {
	rand.Seed(time.Now().UnixNano())
	for _, header := range trolleyHeaders {
		if rand.Float32() < 0.5 { // 50% chance to add each header
			w.Header().Set(header.Key, header.Value)
		}
	}
}
