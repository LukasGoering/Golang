package main

import (
	"fmt"
	"strings"
)

func findIan() {
	// This function checks if a user input fulfills the following conditions:
	// 1. The string starts with an i.
	// 2. The string ends with an n.
	// 3. The string contains the letter a.

	// Initialize the string variable
	var input string
	// Ask the user for input
	fmt.Print("Please enter a string: ")
	fmt.Scan(&input)
	// Make the input case insensitive
	input = strings.ToLower(input)
	// Check the conditions for the string
	if strings.HasPrefix(input, "i") &&
		strings.HasSuffix(input, "n") &&
		strings.Contains(input, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not found!")
	}
}
