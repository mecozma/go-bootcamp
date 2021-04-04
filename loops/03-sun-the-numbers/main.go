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
// EXERCISE: Sum the Numbers
//
//  1. By using a loop, sum the numbers between 1 and 10.
//  2. Print the sum.
//
// EXPECTED OUTPUT
//  Sum: 55
// ---------------------------------------------------------

func main() {
	// Get the commandline args.
	a := os.Args
	if l := len(a); l < 2 {
		fmt.Println("Provide an integer.")
		return
	}

	// Convert to integer and check for negative integer.
	if i, err := strconv.Atoi(a[1]); err != nil {
		fmt.Printf("%s is not an integer.\n", a[1])
	} else if i < 0 {
		fmt.Printf("%d is a negative integer. Please pass a positive integer.\n", i)
	} else {

		// Variable to hold the value.
		var s int
		for j := 1; j <= i; j++ {
			s += j
		}
		fmt.Printf("Sum of %d is:%5d\n", i, s)
	}
}
