package main

import (
	"fmt"
	"os"
)

// A program that greets the user using a default value and the ability to take
// an argument from the command line and print it.

func main() {

	name := "there"

	// next step: if more arguments are provided, combine them to one string and print it out.
	//
	//using concatenation:
	// if len(os.Args) > 1 {
	// 	name = ""
	// 	for _, value := range os.Args {
	// 		name += " " + value
	//
	// 	}
	if len(os.Args) > 1 {

		// now using Sprint:
		name = ""
		for _, value := range os.Args {
			name = fmt.Sprint(name, " ", value)

		}

	}
	fmt.Printf("Hi %v!", name)
}
