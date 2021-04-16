package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {

	// Clear the screen.
	screen.Clear()

	//Print the clock.
	for {
		screen.MoveTopLeft()
		// Get the current time.
		now := time.Now()
		h := now.Hour()
		m := now.Minute()
		s := now.Second()
		ss := now.Nanosecond() / 1e8

		// Clock array.
		clock := [...]placeholder{
			digits[h/10], digits[h%10],
			separator,
			digits[m/10], digits[m%10],
			separator,
			digits[s/10], digits[s%10],
			dot,
			digits[ss],
		}
		for line := range clock[0] {
			for index, digit := range clock {
				blink := clock[index][line]
				if (digit == separator || digit == dot) && s%2 == 0 {
					blink = "   "
				}
				fmt.Print(blink, "  ")
			}
			fmt.Println()
		}
		fmt.Println()

		time.Sleep(time.Second / 10)

	}
}
