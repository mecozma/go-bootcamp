package main

import (
	"fmt"
	"time"
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
	}

	// array to hold all the digits.
	digits := [...]placeholder{zero, one, two, three, four, five, six, seven, eight, nine}

	// Get the current time.
	t := time.Now()

	h := t.Hour()
	m := t.Second()
	s := t.Second()

	// Clock array.
	clock := [...]placeholder{
		digits[h/10], digits[h%10],
		separator,
		digits[m/10], digits[m%10],
		separator,
		digits[s/10], digits[s%10],
	}

	// Print the digits
	// Clear the screen.
	fmt.Print("\033[2J")
	for {
		fmt.Print("\033[H")
		for v := range clock[0] {
			for j := range clock {
				//bl := clock[j][v]
				//if digit == separator && s%2 == 0 {
				//	bl = "   "
				//}
				fmt.Print(clock[j][v], " ")
			}
			fmt.Println()
		}
		fmt.Println()
		time.Sleep(time.Second)
	}

}
