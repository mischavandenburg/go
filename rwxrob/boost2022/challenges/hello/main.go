package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Hello World!")

	// more ways to print
	// This logs and reirects to stderr
	log.Println("Hello World, using log")
	Foo()
}
