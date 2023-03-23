package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello World!")

	// more ways to print
	// This logs and reirects to stderr
	log.Println("Hello World, using log")
	log.Printf("Hello Log World! (PID: %v)", os.Getpid())
	Foo()
}
