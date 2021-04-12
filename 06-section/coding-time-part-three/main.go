package main

import "fmt"

const (
	winter = 1
	summer = 3
	yearly = winter + summer
)

func main() {
	books := [yearly]string{
		"Kafka's Revenge",
		"Stay Golden",
		"Kafka's Revenge Second edition",
	}
	fmt.Printf("books  : %#v\n", books)

	var (
		wBooks [winter]string
		sBooks [summer]string
	)

	wBooks[0] = books[0]
	// sBooks[0] = books[1]
	// sBooks[1] = books[2]
	// sBooks[2] = books[3]

	for i := range sBooks {
		sBooks[i] = books[i+1]
	}
	fmt.Printf("\nwinter : %#v\n", wBooks)
	fmt.Printf("\nsummer : %#v\n", sBooks)

	var published [len(books)]bool
	published[0] = true
	published[len(books)-1] = true

	for i, ok := range published {
		if ok {
			fmt.Printf("+  %s\n", books[i])
		}
	}
}
