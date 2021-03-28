package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

const (
	feetInMeeters float64 = 0.3048
	feetInIards           = feetInMeeters / 0.9144
)

func main() {
	arg := os.Args[1]
	feet, _ := strconv.ParseFloat(arg, 64)
	meters := feet * feetInMeeters
	yards := math.Round(feet * feetInIards)
	fmt.Printf("%g feet is %g meters.\n", feet, meters)
	fmt.Printf("%g feet is %g yards.\n", feet, yards)
}
