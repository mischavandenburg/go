package main

import (
	"fmt"
	"os"
)

// A program that greets the user using a default value and the ability to take
// an argument from the command line and print it.

func main() {

	name := "there"

	if len(os.Args) > 1 {
		name = os.Args[1]
		fmt.Printf("Hi %v!", name)
		os.Exit(0)
	} else {
		fmt.Printf("Hi %v!", name)
	}
}
