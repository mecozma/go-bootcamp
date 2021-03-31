package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("please tell me a value in feet.")
		return
	}
	arg := os.Args[1]

	if feet, err := strconv.ParseFloat(arg, 64); err != nil {
		fmt.Printf("error: '%s' is not a number.\n", arg)
	} else {
		meters := feet * 0.3048
		fmt.Printf("%g feet is %g meters.\n", feet, meters)
	}

}
