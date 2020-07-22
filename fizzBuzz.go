/*
Author: Jasmin A. Smith
Date: 06/04/2020
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Defining future varaibles before hand
	var response string
	var numberOfTimes int
	var count int = 0

	// Asks the user how many fizz buzzs they want
	fmt.Print("How many fizzing and buzzing units fo you need in your life? ")
	// Takes user input
	fmt.Scan(&response)
	// Converts string to int
	numberOfTimes, _ = strconv.Atoi(response)
	// Prints out fizz buzzes for desired amount of time
	for i := 0; count < numberOfTimes; i++ {
		// Prints out fizz for numbers divisible by 3
		if i%3 == 0 && i%5 != 0 {
			fmt.Println("fizz")
			// keeps track of fizz buzz
			count++
			// Prints out buzz for numbers divisible by 5
		} else if i%3 != 0 && i%5 == 0 {
			fmt.Println("buzz")
			// keeps track of fizz buzz
			count++
			// Prints out fizz buzz for numbers divisible by 3 and 5 that is not 0
		} else if i%3 == 0 && i%5 == 0 && i != 0 {
			fmt.Println("fizz buzz")
			// keeps track of fizz buzz
			count++
			// Prints out numbers that are not divisible by 3 or 5
		} else {
			fmt.Println(i)
		}
	}
	fmt.Println("TRADITION!!")
}
