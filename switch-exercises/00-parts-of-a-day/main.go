package main

import (
	"fmt"
	"time"
)

func main() {
	switch h := time.Now().Hour(); {
	case h > 5 && h < 12:
		fmt.Println("good morning!")
	case h > 12 && h < 18:
		fmt.Println("good afternoon!")
	case h > 18 && h < 22:
		fmt.Println("good evening!")
	default:
		fmt.Println("night!")
	}
}
