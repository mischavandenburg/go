package main

import (
	"fmt"
	"time"

	"github.com/mischavandenburg/go/rwxrob/boost2022/challenges/termcolors"
)

func main() {
	for {
		fmt.Print(termcolors.Red + "nyan" + termcolors.Reset)
		time.Sleep(100 * time.Millisecond)
	}
}
