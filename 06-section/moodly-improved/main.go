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
	"math/rand"
	"os"
	"strings"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Moodly #2
//
//   This challenge is based on the previous Moodly challenge.
//
//   1. Display the usage if the username or the mood is missing
//
//   2. Change the moods array to a multi-dimensional array.
//
//      So, create two inner arrays:
//        1. One for positive moods
//        2. Another one for negative moods
//
//   4. Randomly select and print one of the mood messages depending
//      on the given mood command-line argument.
//
// EXPECTED OUTPUT
//
//   go run main.go
//     [your name] [positive|negative]
//
//   go run main.go Socrates
//     [your name] [positive|negative]
//
//   go run main.go Socrates positive
//     Socrates feels good 👍
//
//   go run main.go Socrates positive
//     Socrates feels happy 😀
//
//   go run main.go Socrates positive
//     Socrates feels awesome 😎
//
//   go run main.go Socrates negative
//     Socrates feels bad 👎
//
//   go run main.go Socrates negative
//     Socrates feels sad 😞
//
//   go run main.go Socrates negative
//     Socrates feels terrible 😩
// ---------------------------------------------------------

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("[your name] [positive|negative]")
		return
	}

	name := args[0]
	mood := strings.ToLower(args[1])

	moods := [2][3]string{
		{"happy 😀", "good 👍", "awesome 😎"},
		{"sad 😞", "bad 👎", "terrible 😩"},
	}

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(moods))

	switch mood {
	case "positive":
		fmt.Printf("%s feels %s\n", name, moods[0][n])
	case "negative":
		fmt.Printf("%s feels %s\n", name, moods[1][n])
	}

}
