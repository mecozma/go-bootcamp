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
// EXERCISE: Word Finder
//
//   Your goal is to search for the words inside the corpus.
//
//   Note: This exercise is similar to the previous word finder program:
//   https://github.com/inancgumus/learngo/tree/master/13-loops/10-word-finder-labeled-switch
//
//   1. Get the search query from the command-line (it can be multiple words)
//
//   2. Filter these words, do not search for them:
//      and, or, was, the, since, very
//
//      To do this, use an array for the filtered words.
//
//   3. Print the words found.
//
// RESTRICTION
//   + The search and the filtering should be case insensitive
//
// HINT
//   + strings.Fields function converts a given string to a slice.
//
//     You can find its example in the word finder program that I've mentioned
//     above.
//
// EXPECTED OUTPUT
//   go run main.go
//     Please give me a word to search.
//
//   go run main.go and was
//
//   go run main.go AND WAS
//
//   go run main.go cat beginning
//     #2 : "cat"
//     #11: "beginning"
//
//   go run main.go Cat Beginning
//     #2 : "cat"
//     #11: "beginning"
// ---------------------------------------------------------

const corpus = "lazy cat jumps again and again and again since the beginning this was very important"

func main() {
	// Get the search query from the command-line.
	a := os.Args[1:]
	if l := len(a); l == 0 {
		fmt.Println("Please give me a word to search.")
		return
	}
	ignoreVords := [6]string{"and", "or", "was", "the", "since", "very"}
	words := strings.Fields(corpus)
	// Filter the words and print the results.
	// Iterate over the query words.
	for _, v := range a {
		low := strings.ToLower(v)
		if low == ignoreVords[0] || low == ignoreVords[1] || low == ignoreVords[2] ||
			low == ignoreVords[3] || low == ignoreVords[4] || low == ignoreVords[5] {
			continue
		} else {
			for idx, word := range words {
				if low == word {
					fmt.Printf("#%d : %q\n", idx+1, word)
				}
			}

		}
	}

}
