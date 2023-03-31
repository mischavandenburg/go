package main

import (
	"fmt"
	"os"
)

func output() string {
	// fmt.Printf("%v\n%T", os.Args, os.Args)
	var buf string
	for n, val := range os.Args {
		buf += fmt.Sprintf("$%v  --> %q\n", n, val)
	}
	return buf
}

func main() {
	fmt.Print(output())
}
