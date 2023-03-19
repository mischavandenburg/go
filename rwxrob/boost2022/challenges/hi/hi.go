package main

import (
	"fmt"
)

// A program that greets the user using a default value and the ability to take
// an argument from the command line and print it.

func main() {

	name := "there"

	//	os.Args[1]
	fmt.Printf("Hi %v!", name)
}
