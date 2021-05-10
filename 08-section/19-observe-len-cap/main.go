package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Observe the length and capacity
//
//  Follow the instructions inside the code below to
//  gain more intuition about the length and capacity of a slice.
//
// ---------------------------------------------------------

func main() {
	// --- #1 ---
	// 1. create a new slice named: games
	//
	//var games []string
	// 2. print the length and capacity of the games slice
	//
	//fmt.Printf("length: %d		capacity: %d\n", len(games), cap(games))
	// 3. comment out the games slice
	//    then declare it as an empty slice
	//
	//	emptyGames := []string{}
	// 4. print the length and capacity of the games slice
	//
	//	fmt.Printf("length: %d		capacity: %d\n", len(emptyGames), cap(emptyGames))

	// 5. append the elements: "pacman", "mario", "tetris", "doom"
	//
	//	emptyGames = append(emptyGames, "pacman", "mario", "tetris", "doom")
	// 6. print the length and capacity of the games slice
	//
	//	fmt.Printf("length: %d		capacity: %d\n", len(emptyGames), cap(emptyGames))
	// 7. comment out everything
	//
	// 8. declare the games slice again using a slice literal
	//    (use the same elements from step 3)
	literalGames := []string{"pacman", "mario", "tetris", "doom"}

	// --- #2 ---
	// 1. use a loop from 0 to 4 to slice the games slice, element by element.
	//
	// 2. print its length and capacity along the way (in the loop).

	fmt.Println()
	for i := range literalGames {
		el := literalGames[0:i]
		fmt.Printf("games[:%d]'s len: %d cap: %d\n", i, len(el), cap(el))
	}

	// --- #3 ---
	// 1. slice the games slice up to zero element
	//    (save the result to a new slice named: "zero")
	//
	zero := literalGames[:0]
	// 2. print the games and the new slice's len and cap
	//
	fmt.Printf("literalGames length: %d		capacity: %d\n", len(literalGames), cap(literalGames))
	fmt.Printf("zero length: %d		capacity: %d\n", len(zero), cap(zero))
	// 3. append a new element to the new slice
	//
	zero = append(zero, "zero")
	// 4. print the new slice's lens and caps
	//
	fmt.Printf("zero length: %d		capacity: %d\n", len(zero), cap(zero))
	// 5. repeat the last two steps 5 times (use a loop)
	//
	fmt.Println()
	for i := 0; i < 5; i++ {
		zero = append(zero, "zero")
		fmt.Printf("zero length: %d		capacity: %d\n", len(zero), cap(zero))
	}
	// 6. notice the growth of the capacity after the 5th append
	//
	// Use this slice's elements to append to the new slice:
	fmt.Println()
	elems := []string{"ultima", "dagger", "pong", "coldspot", "zetra"}

	zero = append(zero, elems...)
	fmt.Printf("literalGames's     len: %d cap: %d\n", len(literalGames), cap(literalGames))
	fmt.Printf("zero's      len: %d cap: %d\n", len(zero), cap(zero))

	for _, v := range elems {
		zero = append(zero, v)
		fmt.Printf("zero's      len: %d cap: %d\n", len(zero), cap(zero))
	}

	// --- #4 ---
	// using a range loop, slice the zero slice element by element,
	// and print its length and capacity along the way.
	//
	// observe that, the range loop only loops for the length, not the cap.
	fmt.Println()

	for n := range zero {
		s := zero[:n]
		fmt.Printf("zero[:%d]'s  len: %d cap: %d\n", n, len(s), cap(s))
	}

	// --- #5 ---
	// 1. do the 3rd step above again but this time, start by slicing
	//    the zero slice up to its capacity (use the cap function).
	//
	// 2. print the elements of the zero slice in the loop.
	fmt.Println()
	zero = zero[:cap(zero)]
	for i := range zero {
		s := zero[:i]
		fmt.Printf("zero[:%d]'s  len: %d cap: %d - %q\n", i, len(s), cap(s), s)
	}

	// --- #6 ---
	// 1. change the one of the elements of the zero slice
	zero[0] = "replaced element"
	//
	// 2. change the same element of the games slice
	//
	literalGames[0] = "games replaced element"
	// 3. print the games and the zero slices
	//
	// 4. observe that they don't have the same backing array
	fmt.Println()
	// ...
	fmt.Printf("zero  : %q\n", zero)
	fmt.Printf("games : %q\n", literalGames)

	// --- #7 ---
	// try to slice the games slice beyond its capacity
	//	literalGames = literalGames[:cap(literalGames)+1]

}
