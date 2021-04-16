package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	// for keeping things easy to read and type-safe
	type placeholder [5]string

	zero := placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	one := placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	two := placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	three := placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	four := placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	five := placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	six := placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	seven := placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	eight := placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	nine := placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	separator := placeholder{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}

	// array to hold all the digits.
	digits := [...]placeholder{zero, one, two, three, four, five, six, seven, eight, nine, separator}

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

		// Clock array.
		clock := [...]placeholder{
			digits[h/10], digits[h%10],
			separator,
			digits[m/10], digits[m%10],
			separator,
			digits[s/10], digits[s%10],
		}
		for line := range clock[0] {
			for index, digit := range clock {
				blink := clock[index][line]
				if digit == separator && s%2 == 0 {
					blink = "   "
				}
				fmt.Print(blink, "  ")
			}
			fmt.Println()
		}
		fmt.Println()
		time.Sleep(time.Second)

	}
}
