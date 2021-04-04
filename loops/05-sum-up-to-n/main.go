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
// EXERCISE: Sum up to N
//
//  1. Get two numbers from the command-line: min and max
//  2. Convert them to integers (using Atoi)
//  3. By using a loop, sum the numbers between min and max
//
// RESTRICTIONS
//  1. Be sure to handle the errors. So, if a user doesn't
//     pass enough arguments or she passes non-numerics,
//     then warn the user and exit from the program.
//
//  2. Also, check that, min < max.
//
// HINT
//  For converting the numbers, you can use `strconv.Atoi`.
//
// EXPECTED OUTPUT
//  Let's suppose that the user runs it like this:
//
//    go run main.go 1 10
//
//  Then it should print:
//
//    1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 + 10 = 55
// ---------------------------------------------------------

func main() {
	// Get the args from the command line.
	a := os.Args
	// Check if the user passed the right number of arguments.
	if l := len(a); l != 3 {
		fmt.Println("ussage: go run main.go [min int] [max int]")
		return
	}

	// Convert the arguments to integers.
	if min, minErr := strconv.Atoi(a[1]); minErr != nil {
		fmt.Printf("%d is not an integer\n", min)
	} else if max, maxErr := strconv.Atoi(a[2]); maxErr != nil {
		fmt.Printf("%d is not an integer\n", min)
	} else if min > max {
		fmt.Printf("%d should be smaller than %d\n", min, max)
	} else {
		var sum int
		for i := min; i <= max; i++ {
			sum += i
			fmt.Print(i)
			if i < max {
				fmt.Print(" + ")
			}
		}
		fmt.Printf(" = %d\n", sum)
	}

}
