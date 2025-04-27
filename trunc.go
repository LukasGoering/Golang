package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func trunc() {
	// Initialize a scanner to read from standard input
	scanner := bufio.NewScanner(os.Stdin)

	// Ask the user for input
	fmt.Print("Please enter a float number: ")
	scanner.Scan()
	input := scanner.Text()

	// Convert the input to a float
	floatInt, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid float number.")
		return
	}

	// Truncate the float number to an integer
	intPart := int(floatInt)

	// Print the input and the truncated integer part
	fmt.Print("The number you typed is: ", floatInt, "\n")
	fmt.Print("The truncated integer part is: ", intPart, "\n")
}
