package main

import (
	"fmt"
	"sort"
)

func slice() {
	/* This program starts by creating a slice of integers of length 3.
	It then prompts the user to enter integers, which are added to the slice.
	Each time the user enters an integer, the slice is sorted and printed.
	The program continues until the user types 'X' to quit.
	*/

	// Create an empty slice of integers with length 3 and capacity 10
	initialLength := 3
	initialCapacity := 10
	intSlice := make([]int, initialLength, initialCapacity)
	// Set counter to determine if all initial elements are filled
	currentIndex := 0
	for {
		// Prompt user for input
		fmt.Print("Please enter an integer (or type 'X' to quit): ")
		var input string
		fmt.Scan(&input)

		// Check if the user wants to quit
		if input == "X" || input == "x" {
			break
		}

		// Convert the input to an integer
		/*  Use Sscanf to parse the input to an integer
		    The integer gets stored in num
		    If err is not nil, it means the input was not a valid integer
		*/
		var num int
		_, err := fmt.Sscanf(input, "%d", &num)
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer.")
			continue
		}
		// Replace initial zeros or append to the slice
		if currentIndex < initialLength {
			switch {
			case intSlice[0] == 0:
				intSlice[0] = num
			case intSlice[1] == 0:
				intSlice[1] = num
			case intSlice[2] == 0:
				intSlice[2] = num
			default:
				fmt.Println("All initial elements are filled. Logical error.")
			}
		} else {
			// Append new elements if beyond initial length
			intSlice = append(intSlice, num)
		}
		currentIndex++

		// Sort the slice in ascending order, and print it
		sort.Ints(intSlice)
		fmt.Println("Sorted slice:", intSlice)
	}
}
