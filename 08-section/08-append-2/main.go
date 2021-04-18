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
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Append #2
//
//  1. Create the following nil slices:
//     + Pizza toppings
//     + Departure times
//     + Student graduation years
//     + On/off states of lights in a room
//
//  2. Append them some elements (use your creativity!)
//
//  3. Print all the slices
//
//
// EXPECTED OUTPUT
// (Your output may change, mine is like so:)
//
//  pizza       : [pepperoni onions extra cheese]
//
//  graduations : [1998 2005 2018]
//
//  departures  : [2019-01-28 15:09:31.294594 +0300 +03 m=+0.000325020
//  2019-01-29 15:09:31.294594 +0300 +03 m=+86400.000325020
//  2019-01-30 15:09:31.294594 +0300 +03 m=+172800.000325020]
//
//  lights      : [true false true]
//
//
// HINTS
//  + For departure times, use the time.Time type. Check its documentation.
//
//      now := time.Now()     -> Gives you the current time
//      now.Add(time.Hour*24) -> Gives you a time.Time 24 hours after `now`
//
//  + For graduation years, you can use the int type
// ---------------------------------------------------------

func main() {
	var (
		pizzaToppings     = []string{}
		departureTimes    = []time.Time{}
		studentGraduation = []int{}
		lights            = []bool{}
	)
	pizzaToppings = append(pizzaToppings, "red onions", "olives", "mushrums")
	now := time.Now()
	t1 := now.Add(time.Hour * 1)
	t2 := now.Add(time.Hour * 2)
	t3 := now.Add(time.Hour * 3)
	t4 := now.Add(time.Hour * 4)
	departureTimes = append(departureTimes, t1, t2, t3, t4)
	studentGraduation = append(studentGraduation, 1885, 1999, 5345)
	lights = append(lights, true, false, false)

	fmt.Println("pizza: ", pizzaToppings)
	fmt.Println("departures:", departureTimes)
	fmt.Println("graduations: ", studentGraduation)
	fmt.Println("lights: ", lights)
}
