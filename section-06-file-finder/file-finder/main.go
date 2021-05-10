package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// get the command line arguments.
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Provide a directory")
		return
	}

	// read the files in the provided directory.
	files, err := ioutil.ReadDir(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	// calculate the backing aray of the slice.
	var total int

	for _, file := range files {
		if file.Size() == 0 {
			total += len(file.Name()) + 1 // +1 for each new line.
		}
	}
	fmt.Printf("Total required space: %d bytes.\n", total)

	names := make([]byte, 0, total)

	for _, file := range files {
		s := file.Size()
		n := file.Name()
		if s == 0 {
			names = append(names, n...)
			names = append(names, '\n')
		}
		// write to file.
		ioutil.WriteFile("empty-files.txt", names, 0644)

	}
	fmt.Printf("%s", names)
}
