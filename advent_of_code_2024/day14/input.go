package main

import (
	"fmt"
	"os"
)

func PromptUserForCookie() string {
	// Prompt user for Cookie.
	fmt.Println("Enter your session cookie.")

	var cookie string

	fmt.Scanln(&cookie)

	return cookie
}

func GetCookieFromEnvVar() string {
	return os.Getenv("AOC2024_SESSION_COOKIE")
}
