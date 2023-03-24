package main

import (
	"fmt"
	"os"
)

// variadic function can take 0 or more arguments
func SayHi(args ...string) {

	// fmt.Println("Hi " + name + "!")
	name := "there"
	if len(args) > 0 {
		name = args[0]
	}
	fmt.Printf("Hi %v!\n", name)
}

// the first argument of an executable is always the path to the executable
// therefore we want evertyhing except the first argument, so [1:]
func main() {
	SayHi(os.Args[1:]...)
}
