package main

import (
	"math/rand"
	"time"
)

// List of comical responses
var comicalResponses = []string{
	"Oops! This endpoint seems to have run away from home. Please try again later.",
	"You've found the secret API endpoint! Just kidding, it's not so secret after all.",
	"This data is so classified, even we forgot what it was supposed to be. Proceed with caution.",
	"Congratulations! You've discovered a hidden feature: Absolutely Nothing™.",
	"This endpoint is currently on vacation. Please leave a message after the beep.",
	"If you were hoping for sensitive data, you'd better hope harder. This ain't it, chief.",
	"Fun fact: This API endpoint was coded in an ancient language only known to the wise ones.",
	"Error 418: I'm a teapot. Just kidding, but seriously, this endpoint does nothing useful.",
}

// Function to get a random comical response
func getRandomComicalResponse() string {
	rand.Seed(time.Now().UnixNano())
	return comicalResponses[rand.Intn(len(comicalResponses))]
}
