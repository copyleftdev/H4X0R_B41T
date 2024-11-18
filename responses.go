package main

import (
	"math/rand"
	"time"
)

// List of comical responses with added false positives that look legit
var comicalResponses = []string{
	"Oops! This endpoint seems to have run away from home. Please try again later.",
	"You've found the secret API endpoint! Just kidding, it's not so secret after all.",
	"This data is so classified, even we forgot what it was supposed to be. Proceed with caution.",
	"Congratulations! You've discovered a hidden feature: Absolutely Nothing™.",
	"This endpoint is currently on vacation. Please leave a message after the beep.",
	"If you were hoping for sensitive data, you'd better hope harder. This ain't it, chief.",
	"Fun fact: This API endpoint was coded in an ancient language only known to the wise ones.",
	"Error 418: I'm a teapot. Just kidding, but seriously, this endpoint does nothing useful.",
	"Warning: This API is protected by quantum encryption. Or is it?",
	"Access Denied. But feel free to try again for no apparent reason.",
	"System Error: Brain not found. Please reboot the universe.",
	"Unauthorized access detected. Security ninjas have been deployed.",
	"Congratulations! You've unlocked the Easter egg: The egg is empty.",
	"This endpoint self-destructed out of boredom. Try another one.",
	"Error 404: The meaning of life, the universe, and everything is still unknown.",
	"You've triggered the alarm! Just kidding, or am I?",
	"This endpoint is currently being guarded by a very angry rubber duck.",
	"Nice try, hacker! But this API is powered by sarcasm and irony.",
	"Warning: This API is cursed. Proceed at your own risk.",
	"This endpoint requires a unicorn-level clearance. Do you have one?",
	"Error: You seem to be using a browser from the Stone Age.",
	"Congratulations! You've found the legendary NullPointerException.",
	"Did you really think there was something valuable here? That's adorable.",
	"Accessing this endpoint will not make you rich or famous. Trust me.",
	"Data leakage detected! Just kidding, the only thing leaking is sarcasm.",
	"System update required: Please install common sense v2.0.",
	"Hey there, hacker! Need a guide to real vulnerabilities?",
	"Fun fact: This API drinks more coffee than you do.",
	"Alert: Hacker detected! Deploying virtual pies in 3... 2... 1...",
	"This endpoint is powered by black holes. No data will ever escape.",
	"Error 666: The demons of the API are not impressed by your intrusion.",
	"Hey, script kiddie! Why don't you learn some real skills?",
	"Nice try, but this API is protected by Chuck Norris' firewall.",
	"Critical Error: The hamster powering this API just fell off the wheel.",
	"Congratulations! You've successfully wasted a few seconds of your life.",
	"This API is under the witness protection program. Keep it secret.",
	"System overload: Too many bad requests detected. Take a break, hacker.",
	"Fun fact: Trying harder won’t make this endpoint less useless.",
	"Nice attempt! But the treasure you seek is in another API castle.",
	"This API endpoint is guarded by a sarcastic AI. Welcome to the club.",
	"Error: The developer who built this forgot to give it a purpose.",
	"This is not the endpoint you are looking for. Move along, Jedi.",
	"Access granted! Just kidding, you wish.",
	"Detected vulnerability: Your ego. Patch it immediately.",
	"This API was written by a wizard. Expect the unexpected.",
	"Critical warning: You should really stop poking around here.",
	"You've unlocked... absolutely nothing! Better luck next time.",
	"Error: Attempting to hack this API has downgraded your karma.",
	"Alert: You just made this API laugh. It needed that.",
	"Security alert: A bucket of sarcasm has been deployed.",
	"Nice guess! But all the real data is stored on a USB drive in Atlantis.",
	"Attention: The hacker gods are not impressed by your skills.",
	"Oops! This endpoint is powered by good intentions and bad coding.",
	// Legit-Looking False Positives
	"Database connection error: Unable to establish a secure link. Please try again.",
	"Access token expired. Request a new token from the /auth endpoint.",
	"Configuration error: Missing API key. Please check your request headers.",
	"Warning: Deprecated API version detected. Upgrade to version 2.1 for security enhancements.",
	"Notice: Your IP has been logged for security auditing purposes.",
	"Authentication failure: Incorrect credentials provided.",
	"Session expired. Please re-authenticate to continue.",
	"Error: Failed to retrieve data from the primary database. Attempting failover...",
	"Invalid API key. Contact support if you believe this is an error.",
	"Security alert: Multiple failed access attempts detected.",
	"Data access violation: Insufficient permissions for this resource.",
	"Notice: This endpoint is only available to premium users.",
	"Error: API rate limit exceeded. Try again in 60 seconds.",
	"Internal server error: Please try your request again later.",
	"Warning: Unusual activity detected from your IP address.",
	"Notice: Maintenance mode active. API access restricted until further notice.",
}

// Function to get a random comical response
func getRandomComicalResponse() string {
	rand.Seed(time.Now().UnixNano())
	return comicalResponses[rand.Intn(len(comicalResponses))]
}
