package main

import (
	"fmt"
	"time"

	"github.com/mischavandenburg/go/rwxrob/boost2022/challenges/termcolors"
	C "github.com/mischavandenburg/go/rwxrob/boost2022/challenges/termcolors"
)

func main() {
	fmt.Print(C.CurOff)
	for {
		fmt.Print(termcolors.Rand() + "nyan" + termcolors.Reset)
		time.Sleep(10 * time.Millisecond)
	}
}
