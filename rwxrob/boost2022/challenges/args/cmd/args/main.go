package main

import (
	"fmt"
	"os"

	"github.com/mischavandenburg/go/rwxrob/boost2022/challenges/args/internal"
)

func main() {
	out := internal.Output(os.Args...)
	fmt.Print(out)
}
