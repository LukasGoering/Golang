package main

import (
	"bufio"
	"fmt"
	"os"
)

// Scanner for reading input
// os for file operations

func readNames() {
	// This program will read names from a text file and store them in a structure.
	// It will then print the names to the console.

	// Create the structure to hold the data
	type Person struct {
		firstName string
		lastName  string
	}

	// Initialize a slice to hold the Person structures
	var people []Person

	// Initialize a scanner to read from standard input
	scanner := bufio.NewScanner(os.Stdin)

	// Prompt the user for the name of the text file to read
	fmt.Printf("Please enter the name of the text file to read: ")
	scanner.Scan()
	fileName := scanner.Text()

	// Open the file for reading
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	// Close the text file after main function is done
	defer file.Close()

	// Create a new scanner for the file
	file_scanner := bufio.NewScanner(file)

	// Read the file line by line
	for file_scanner.Scan() {
		// Save the currently scanned line to a string variable
		line := file_scanner.Text()
		// Split the line into first and last names
		var firstName, lastName string
		_, err := fmt.Sscanf(line, "%s %s", &firstName, &lastName)

		// Check if the line was read correctly
		if err != nil {
			fmt.Println("Error reading line:", err)
			continue
		}

		// Create a new Person structure and append it to the slice
		person := Person{firstName: firstName, lastName: lastName}
		people = append(people, person)
	}

	// Print the names to the console
	for _, person := range people {
		fmt.Printf("%s %s\n", person.firstName, person.lastName)
	}
}
