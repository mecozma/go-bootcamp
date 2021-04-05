package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var (
	maxTurns = 5
	usage    = `Welcome to the Lucky Number Game!


The program will pick %d random  numbers.
Your mission is to guess one of those numbers.

The greater your number is, the greater it gets.

Wanna play?`
)

func main() {
	// Seed the random package with time.
	rand.Seed(time.Now().UnixNano())

	// Get the command line arguments.
	a := os.Args
	if l := len(a); l < 3 {
		fmt.Printf(usage, maxTurns)
		return
	}

	// Convert the guess number to int and chack the errors.
	guess, err := strconv.Atoi(a[1])
	secondGuess, secondErr := strconv.Atoi(a[2])
	if err != nil && secondErr != nil {
		fmt.Println("Invalid input!")
		return
		//Check if the integer is positive.
	} else if guess < 0 || secondGuess < 0 {
		fmt.Println("Pick a positive number.")
		return
	} else {

		// Game dificulty logic.
		var difficulty = 0
		if guess < 10 && secondGuess < 10 {
			difficulty = 10
		} else {
			difficulty = int(math.Max(float64(guess), float64(secondGuess)))
		}
		// Loop for 5 times.
		for turn := 0; turn < maxTurns; turn++ {
			n := rand.Intn(difficulty + 1)
			fmt.Printf("%5d     ", turn)
			if n == guess || n == secondGuess {
				switch rand.Intn(3) {
				case 0:
					fmt.Println("You Won!")
				case 1:
					fmt.Println("You nailed it!")
				case 2:
					fmt.Println("Fair Play! You won!")
				}

				if turn == 0 {
					fmt.Println("++ bonus ++ ")
				}
				fmt.Printf("Your Lucky Number is: %5d\n", n)
				return
			}
		}
	}
	switch rand.Intn(3) {
	case 0:
		fmt.Println("Try again!")
	case 1:
		fmt.Println("Not this time!")
	case 2:
		fmt.Println("That's embarrassing!")
	}
}
