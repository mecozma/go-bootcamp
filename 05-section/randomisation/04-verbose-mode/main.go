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
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Verbose Mode
//
//  When the player runs the game like this:
//
//    go run main.go -v 4
//
//  Display each generated random number:

//    1 3 4 🎉  YOU WIN!
//
//  In this example, computer picks 1, 3, and 4. And the
//  player wins.
//
// HINT
//  You need to get and interpret the command-line arguments.
// ---------------------------------------------------------
const (
	maxTurns = 5 // less is more difficult
	usage    = `Welcome to the Lucky Number Game! 🍀
The program will pick %d random numbers.
Your mission is to guess one of those numbers.
The greater your number is, harder it gets.
Wanna play?
(Provide -v flag to see the picked numbers.)
`
)

func main() {
	// Get the command line arguments.
	a := os.Args
	if l := len(a); l < 2 {
		fmt.Printf(usage, maxTurns)
		return
	}

	// Check if the verbose flag is passed.
	var verbose bool
	if a[1] == "-v" {
		verbose = true
	}

	// Check if the passed argument is an int and handle the error.
	guess, err := strconv.Atoi(a[len(a)-1])
	if err != nil {
		fmt.Println("ussage: go run main.go [-v] int")
		return
	}

	// Check if the passed argument is less negative.
	if guess <= 0 {
		fmt.Println("Please pass a positive integer.")
		return
	}

	// Game logic.
	for turn := 1; turn <= maxTurns; turn++ {
		n := rand.Intn(guess) + 1

		if verbose {
			fmt.Printf("%4d", n)
		}

		if n == guess {
			if turn == 1 {
				fmt.Println("First time winner.")
			} else {
				fmt.Println("You won!")
			}
			return
		}
	}
	fmt.Println("You lost...Try again?")

}
