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
// EXERCISE: Only Evens
//
//  1. Extend the "Sum up to N" exercise
//  2. Sum only the even numbers
//
// RESTRICTIONS
//  Skip odd numbers using the `continue` statement
//
// EXPECTED OUTPUT
//  Let's suppose that the user runs it like this:
//
//    go run main.go 1 10
//
//  Then it should print:
//
//    2 + 4 + 6 + 8 + 10 = 30
// ---------------------------------------------------------

func main() {
	// Get command line arguments
	a := os.Args
	if l := len(a); l != 3 {
		fmt.Println("Please pass two command line integer arguments.")
		return
	}

	// Convert args to integers.
	min, errMin := strconv.Atoi(a[1])
	max, errMax := strconv.Atoi(a[2])

	// Check if the passed arguments are integers.
	if errMin != nil || errMax != nil || min >= max {
		fmt.Println("Please make sure the arguments are integers.")
		return
	}

	var sum int

	for i := min; i <= max; i++ {
		if i%2 != 0 {
			continue
		}
		sum += i
		fmt.Print(i)
		if i < max {
			fmt.Print(" + ")
		}
	}
	fmt.Printf(" = %d\n", sum)
}
