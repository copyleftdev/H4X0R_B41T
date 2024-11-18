package main

import (
	"math/rand"
	"net/http"
	"time"
)

// List of trolley headers with some false positives
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
	{"X-Uptime", "42 minutes (and counting)"},
	{"X-Runtime", "Slower than your grandma's dial-up"},
	{"X-Framework", "YoMama.js (Beta)"},
	{"X-AI-Powered", "By a hamster on a wheel"},
	{"X-Coffee-Consumed", "All of it"},
	{"X-Bug-Count", "Infinite"},
	{"X-Fun-Fact", "This API won a staring contest with a database once."},
	{"X-Assistant", "ChatGPT at your service"},
	{"X-Quantum-Leap", "One leap ahead of hackers"},
	{"X-Surprise-Feature", "It's not a bug, it's a feature!"},
	{"X-Y2K-Ready", "Still stuck in the year 2000"},
	{"X-Debug-Mode", "On (or is it?)"},
	{"X-Unicorn-Status", "Galloping in progress"},
	{"X-Meme-Powered", "By the power of Rickroll"},
	{"X-Request-ID", "42 (the answer to everything)"},
	{"X-World-Domination", "Loading..."},
	{"X-Language", "Leetspeak 1337"},
	{"X-Trust-Us", "We’re professionals"},
	{"X-Banana", "For scale"},
	{"X-Under-Construction", "Since 1999"},
	{"X-Magic-Spell", "Wingardium Leviosa!"},
	{"X-Security-Level", "Paranoid++"},
	{"X-Cookie", "Nom nom nom"},
	{"X-Future-Proof", "Tested for Mars missions"},
	{"X-Optimized-For", "Windows ME (Best OS Ever)"},
	{"X-Rickroll", "Never gonna give you up"},
	{"X-Timezone", "Somewhere in the Matrix"},
	{"X-Sarcasm-Mode", "Always on"},
	{"X-Safety", "No refunds if hacked"},
	{"X-Cheat-Code", "Up, Up, Down, Down, Left, Right, B, A"},
	{"X-Maintained-By", "A team of overworked squirrels"},
	{"X-Deprecated", "Since before it was cool"},
	{"X-Chaos-Engine", "Running smoothly"},
	{"X-Version", "∞.0 (The final frontier)"},
	{"X-Creative-Director", "The cat that walked on my keyboard"},
	{"X-Pizza", "Pineapple goes here"},
	{"X-Dystopia", "Welcome to the future"},
	{"X-Timeline", "We're in the darkest one"},
	{"X-Reality", "Suspended until further notice"},
	{"X-Nostalgia", "Powered by floppy disks"},
	{"X-Philosophy", "Existential crisis every hour"},
	{"X-Dad-Joke", "Why did the scarecrow win an award? Because he was outstanding in his field!"},
	{"X-Fail-Safe", "Activate in case of hacker panic"},
	{"X-Documentation", "Written in ancient runes"},
	{"X-Virtual-Reality", "Still virtual, not real"},
	{"X-Bitcoin", "Not enough to make you rich"},
	// False Positives
	{"X-DB-Connection", "jdbc:mysql://127.0.0.1:3306/h4x0r_b41t_db"},
	{"X-Auth-Token", "ExpiredToken12345"},
	{"X-API-Key", "f4k3-4p1-k3y-1337"},
	{"X-Backup-Server", "192.168.1.100 (offline)"},
	{"X-SSH-Port", "2222 (Closed)"},
	{"X-Admin-Contact", "admin@example.com"},
	{"X-Error", "Stack Overflow Detected"},
	{"X-Last-Patched", "Unpatched since 2017"},
	{"X-Vulnerability", "CVE-2023-1234 (Still unfixed)"},
	{"X-Encryption-Key", "Lost and never found"},
	{"X-Password-Reset", "Enabled via /reset-password"},
	{"X-SQL-Mode", "Loose (no strict checks)"},
	{"X-DNS", "8.8.8.8"},
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
