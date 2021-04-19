// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Housing Prices
//
//  We have received housing prices. Your task is to load the data into
//  appropriately typed slices then print them.
//
//  1. Check out the expected output
//
//
//  2. Check out the code below
//
//     1. header   : stores the column headers
//     2. data     : stores the real data related to each column
//     3. separator: you will use it to separate the data
//
//
//  3. Parse the data
//
//     1. Separate it into rows by using the newline character ("\n")
//
//     2. For each row, separate it by using the separator (",")
//
//
//  4. Load the data into distinct slices
//
//     1. Load the locations into a []string
//     2. Load the sizes into []int
//     3. Load the beds into []int
//     4. Load the baths into []int
//     5. Load the prices into []int
//
//
//  5. Print the header
//
//     1. Separate it by using the separator
//
//     2. Print each column 15 character wide ("%-15s")
//
//
//  6. Print the rows from the slices that you've created, line by line
//
//
// EXPECTED OUTPUT
//
//  Location       Size           Beds           Baths          Price
//  ===========================================================================
//  New York       100            2              1              100000
//  New York       150            3              2              200000
//  Paris          200            4              3              400000
//  Istanbul       500            10             5              1000000
//
//
// HINTS
//
//  + strings.Split function can separate a string into a []string
//
// ---------------------------------------------------------

func main() {
	const (
		header = "Location,Size,Beds,Baths,Price"
		data   = `New York,101,2,1,100000
New York,151,3,2,200000
Paris,201,4,3,400000
Istanbul,501,10,5,1000000`

		separator = ","
	)
	h := strings.Split(header, separator)
	d := strings.Split(data, "\n")

	// Load data in distinct slices.
	var (
		locations []string
		sizes     []int
		beds      []int
		baths     []int
		prices    []int
		avgSize   float64
		avgBeds   float64
		avgBaths  float64
		avgPrice  float64
	)

	for _, row := range d {
		columns := strings.Split(row, separator)

		locations = append(locations, columns[0])

		// Variable a is declared and instantiated.
		a, _ := strconv.Atoi(columns[1])
		avgSize += float64(a)
		sizes = append(sizes, a)

		// Variable a value is overwritten.
		a, _ = strconv.Atoi(columns[2])
		avgBeds += float64(a)
		beds = append(beds, a)

		a, _ = strconv.Atoi(columns[3])
		avgBaths += float64(a)
		baths = append(baths, a)

		a, _ = strconv.Atoi(columns[4])
		avgPrice += float64(a)
		prices = append(sizes, a)

	}

	// Print data.
	for _, v := range h {
		fmt.Printf("%-15s", v)
	}

	fmt.Printf("\n%s\n", strings.Repeat("=", 65))

	for i := range d {
		fmt.Printf("%-15s", locations[i])
		fmt.Printf("%-15d", sizes[i])
		fmt.Printf("%-15d", beds[i])
		fmt.Printf("%-15d", baths[i])
		fmt.Printf("%-15d", prices[i])
		fmt.Println()
	}

	fmt.Printf("\n%s\n", strings.Repeat("=", 65))

	fmt.Printf("%-15d", 0)
	fmt.Printf("%-15.2f", avgSize/4)
	fmt.Printf("%-15.2f", avgBeds/4)
	fmt.Printf("%-15.2f", avgBaths/4)
	fmt.Printf("%-15.2f", avgPrice/4)
	fmt.Println()

}
