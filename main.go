package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// findIan searches user input for the name Ian.
	findIan()
	clearInputBuffer()
	// trunc truncates a float number to an integer and prints it.
	trunc()
	clearInputBuffer()
	// makejson creates a JSON object from user input.
	makejson()
	clearInputBuffer()
	// slice adds integers to a slice, sorts it, and prints it.
	slice()
	clearInputBuffer()
	// read reads names from a text file, stores them in structs, and prints them.
	readNames()
	// Use pointers to modify a variable's value.
	var x int = 2
	addOne(&x)
	fmt.Print(x)
}

func clearInputBuffer() {
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n') // Consume the leftover input
}

func addOne(y *int) {
	*y = *y + 1
}
