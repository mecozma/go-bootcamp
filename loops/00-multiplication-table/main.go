package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	//Get the argument passed by the user.
	a := os.Args
	if l := len(a); l < 2 {
		fmt.Println("provide a multiplication factor.")
		return
	}

	//Convert the argument to int.
	mf, _ := strconv.Atoi(a[1])
	fmt.Printf("%5s", "X")
	for i := 0; i <= mf; i++ {
		fmt.Printf("%5d", i)
	}

	fmt.Println()

	for i := 0; i <= mf; i++ {
		fmt.Printf("%5d", i)

		for j := 0; j <= mf; j++ {
			fmt.Printf("%5d", i*j)
		}
		fmt.Println()
	}
}
