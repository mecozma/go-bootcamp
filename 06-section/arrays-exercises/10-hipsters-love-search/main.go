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
// EXERCISE: Hipster's Love Bookstore Search Engine
//
//  Your goal is to allow people to search for books.
//
//  1. Create an array with the following book titles:
//      Kafka's Revenge
//      Stay Golden
//      Everythingship
//      Kafka's Revenge 2nd Edition
//
//  2. Get the search query from the command-line argument
//
//  3. Search for the books in the books array
//
//  4. When the program finds the book, print it.
//  5. Otherwise, print that the book doesn't exist.
//
//  6. Handle the errors.
//
// RESTRICTION:
//   + The search should be case insensitive.
//
// EXPECTED OUTPUT
//   go run main.go
//     Tell me a book title
//
//   go run main.go STAY
//     Search Results:
//     + Stay Golden
//
//   go run main.go sTaY
//     Search Results:
//     + Stay Golden
//
//   go run main.go "Kafka's Revenge"
//     Search Results:
//     + Kafka's Revenge
//     + Kafka's Revenge 2nd Edition
//
//   go run main.go void
//     Search Results:
//     We don't have the book: "void"
//
// HINTS:
//   + To find out whether a string contains another string value, you can use the strings.Contains function.
//   + To convert a string value to lowercase, you can use the strings.ToLower function.
//   + Check out the strings package for more information.
// ---------------------------------------------------------

func main() {
	// Book array.
	books := [...]string{
		"Kafka's Revenge",
		"Stay Golden",
		"Kafka's Revenge 2nd Edition",
	}

	// Get the command line arguments.
	a := os.Args[1:]
	if l := len(a); l < 1 {
		fmt.Println("Tell me a book title.")
		return
	}

	// Book title.
	title := strings.ToLower(a[0])

	fmt.Println("Search Results:")

	// Flag to identify if any books have been found.
	var flag bool

	// Iterate the books.
	for _, book := range books {
		if strings.Contains(strings.ToLower(book), title) {
			fmt.Println("+", book)
			flag = true
		}
	}

	if !flag {
		fmt.Printf("We don't have the book: %q\n", a[0])
	}

}
