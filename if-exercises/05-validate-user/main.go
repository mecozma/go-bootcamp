package main

import (
	"fmt"
	"os"
)

func main() {
	a := os.Args
	l := len(a)
	username := "username"
	password := "password"
	username2 := "username2"
	password2 := "password2"
	usage := "Usage: [username] [password]"
	if l < 3 {
		fmt.Println(usage)
	} else if username == a[1] && password == a[2] ||
		username2 == a[1] && password2 == a[2] {
		fmt.Printf("Access granted for \"%s\"\n", a[1])
	} else {
		fmt.Printf("Access denied for \"%s\"\n", a[1])
	}

}
