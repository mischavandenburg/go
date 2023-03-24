package main

import (
	"os"

	"github.com/mischavandenburg/go/rwxrob/boost2022/challenges/greet"
)

func main() {
	var name string
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	// TODO if name empty prompt for input
	greet.Hi(name)
}
