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
// EXERCISE: Dynamic Difficulty
//
//  Current game picks only 5 numbers (5 turns).
//
//  Make sure that the game adjust its own difficulty
//  depending on the guess number.
//
// RESTRICTION
//  Do not make the game to easy. Only adjust the
//  difficulty if the guess is above 10.
//
// EXPECTED OUTPUT
//  Suppose that the player runs the game like this:
//    go run main.go 5
//
//    Then the computer should pick 5 random numbers.
//
//  Or, if the player runs it like this:
//    go run main.go 25
//
//    Then the computer may pick 11 random numbers
//    instead.
//
//  Or, if the player runs it like this:
//    go run main.go 100
//
//    Then the computer may pick 30 random numbers
//    instead.
//
//  As you can see, greater guess number causes the
//  game to increase the game turns, which in turn
//  adjust the game's difficulty dynamically.
//
//  Because, greater guess number makes it harder to win.
//  But, this way, game's difficulty will be dynamic.
//  It will adjust its own difficulty depending on the
//  guess number.
// ---------------------------------------------------------
const (
	maxTurns int = 5 // less is more difficult
	usage        = `Welcome to the Lucky Number Game! 🍀
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
	for turn := maxTurns + guess/4; turn > 0; turn-- {
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
