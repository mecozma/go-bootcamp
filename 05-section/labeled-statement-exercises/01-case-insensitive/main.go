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
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Case Insensitive Search
//
//  Allow for case-insensitive searching
//
// EXAMPLE
//  Let's say that the user runs the program like this:
//    go run main.go LAZY
//
//  Or like this: go run main.go lAzY
//  Or like this: go run main.go lazy
//
//  For all cases above, the program should find
//  the "lazy" keyword.
// ---------------------------------------------------------

var s string = "the brown fox jumps again and again and again"

func main() {
	words := strings.Fields(s)
	query := os.Args[1:]
queries:
	for _, q := range query {
		// Case insensitive search.
		q = strings.ToLower(q)
	search:
		for i, w := range words {
			switch q {
			case "and", "or", "the":
				break search
			}
			if q == w {
				fmt.Printf("# %-2d: %q\n", i+1, w)
				continue queries
			}

		}
	}
}
