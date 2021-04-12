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
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Moodly
//
//   1. Get username from command-line
//
//   2. Display the usage if the username is missing
//
//   3. Create an array
//     1. Add three positive mood messages
//     2. Add three negative mood messages
//
//   4. Randomly select and print one of the mood messages
//
// EXPECTED OUTPUT
//
//   go run main.go
//     [your name]
//
//   go run main.go Socrates
//     Socrates feels good 👍
//
//   go run main.go Socrates
//     Socrates feels bad 👎
//
//   go run main.go Socrates
//     Socrates feels sad 😞
//
//   go run main.go Socrates
//     Socrates feels happy 😀
//
//   go run main.go Socrates
//     Socrates feels awesome 😎
//
//   go run main.go Socrates
//     Socrates feels terrible 😩
// ---------------------------------------------------------

func main() {

	var m [6]string
	m[0] = "feels good 👍"
	m[1] = "feels bad 👎"
	m[2] = "feels sad 😞"
	m[3] = "feels terrible 😩"
	m[4] = "feels happy 😀"
	m[5] = "feels awesome 😎"

	// get a random number.
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(len(m))
	// get the args from the command line.
	a := os.Args
	// check if there are any parameters passed
	if l := len(a); l < 2 {
		fmt.Println("[your name]")
		return
	} else {
		fmt.Printf("%s %s\n", a[1], m[randomNumber])
	}
}
