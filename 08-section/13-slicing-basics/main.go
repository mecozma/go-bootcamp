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
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Slice the numbers
//
//   We've a string that contains even and odd numbers.
//
//   1. Convert the string to an []int
//
//   2. Print the slice
//
//   3. Slice it for the even numbers and print it (assign it to a new slice variable)
//
//   4. Slice it for the odd numbers and print it (assign it to a new slice variable)
//
//   5. Slice it for the two numbers at the middle
//
//   6. Slice it for the first two numbers
//
//   7. Slice it for the last two numbers (use the len function)
//
//   8. Slice the evens slice for the last one number
//
//   9. Slice the odds slice for the last two numbers
//
//
// EXPECTED OUTPUT
//   go run main.go
//
//     nums        : [2 4 6 1 3 5]
//     evens       : [2 4 6]
//     odds        : [1 3 5]
//     middle      : [6 1]
//     first 2     : [2 4]
//     last 2      : [3 5]
//     evens last 1: [6]
//     odds last 2 : [3 5]
//
//
// NOTE
//
//  You can also use my prettyslice package for printing the slices.
//
//
// HINT
//
//  Find a function in the strings package for splitting the string into []string
//
// ---------------------------------------------------------

func main() {
	// uncomment the declaration below
	data := "2 4 6 1 3 5"
	d := strings.Split(data, " ")
	ints := []int{}

	for _, v := range d {
		a, _ := strconv.Atoi(v)
		ints = append(ints, a)
	}
	fmt.Println(ints)
	// Even numbers.
	even := ints[0:3]
	fmt.Println("evens:", even)

	// Odd numbers.
	odds := ints[3:]
	fmt.Println("odds:", odds)

	// Middle of the ints variable.
	middle := ints[2:4]
	fmt.Println("middle:", middle)

	// First two elements.
	fTwo := ints[0:2]
	fmt.Println("first 2:", fTwo)

	// Last two elemtnts.
	lTwo := ints[len(ints)-2:]
	fmt.Println("last 2:", lTwo)

	// Evens last 1.
	eLast := even[len(even)-1:]
	fmt.Println("evens last 1:", eLast)

	// Odds last 2.
	oLast := odds[len(odds)-2:]
	fmt.Println("odds last 2:", oLast)
}
