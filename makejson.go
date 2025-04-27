package main

import (
	"bufio" // Scanner for reading input
	"encoding/json"
	"fmt"
	"os" // os for file operations
)

func makejson() {
	// This program prompts the user for a name and address, stores them in a map,
	// and then converts the map to a JSON object.

	// Initialize a scanner to read from standard input
	scanner := bufio.NewScanner(os.Stdin)

	// Create a map to store the user input
	person := make(map[string]string)

	// Promt the user for a name and an address
	fmt.Printf("Please enter a name: ")
	scanner.Scan()
	name := scanner.Text()
	person["name"] = name

	fmt.Printf("Please enter an address: ")
	scanner.Scan()
	address := scanner.Text()
	person["address"] = address

	// Create a JSON object from the map
	jsonObject, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	// Print the JSON object as byte code and as a string
	fmt.Println(jsonObject)
	fmt.Println(string(jsonObject))
}
