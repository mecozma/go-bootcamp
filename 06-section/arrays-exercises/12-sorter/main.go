// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Number Sorter
//
//  Your goal is to sort the given numbers from the command-line.
//
//  1. Get the numbers from the command-line.
//
//  2. Create an array and assign the given numbers to that array.
//
//  3. Sort the given numbers and print them.
//
// RESTRICTION
//   + Maximum 5 numbers can be provided
//   + If one of the arguments are not a valid number, skip it
//
// HINTS
//  + You can use the bubble-sort algorithm to sort the numbers.
//    Please watch this: https://youtu.be/nmhjrI-aW5o?t=7
//
//  + When swapping the elements, do not check for the last element.
//
//    Or, you will receive this error:
//    "panic: runtime error: index out of range"
//
// EXPECTED OUTPUT
//   go run main.go
//     Please give me up to 5 numbers.
//
//   go run main.go 6 5 4 3 2 1
//     Sorry. Go arrays are fixed. So, for now, I'm only supporting sorting 5 numbers...
//
//   go run main.go 5 4 3 2 1
//     [1 2 3 4 5]
//
//   go run main.go 5 4 a c 1
//     [0 0 1 4 5]
// ---------------------------------------------------------
func bubbleSort(arr [5]int) {
	var noSwaps bool
	for i := len(arr); i > 0; i-- {
		noSwaps = true
		for j := 0; j < i; j++ {
			if j < len(arr)-1 && arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
				noSwaps = false
			}
		}
		if noSwaps {
			break
		}
	}
	fmt.Printf("The sorted array is: %v\n", arr)
}

func main() {
	// Get the numbers from the command-line.
	a := os.Args[1:]
	if l := len(a); l < 1 {
		fmt.Println("Please give me up to 5 numbers.")
		return
	} else if l > 5 {
		fmt.Println("Sorry. Go arrays are fixed. So, for now, I'm only supporting sorting 5 numbers...")
		return
	}

	// Create an array and assign the given numbers to that array.
	var numbers [5]int
	for i, v := range a {
		if num, err := strconv.Atoi(v); err != nil {
			continue
		} else {
			numbers[i] = num
		}
	}

	// Sort the given numbers and print them.
	bubbleSort(numbers)
}
