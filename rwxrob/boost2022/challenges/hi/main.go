package main

import (
	"fmt"
)

func SayHi(name string) {
	// fmt.Println("Hi " + name + "!")
	fmt.Printf("Hi %v!\n", name)
}

func main() {
	SayHi("Rob")
}
