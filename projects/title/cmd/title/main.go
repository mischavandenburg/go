package main

import (
	"log"
	"os"
	"strings"

	"github.com/mischavandenburg/go/projects/title"
	"github.com/mischavandenburg/go/projects/title/internal"
)

func main() {
	var x string
	var err error
	if len(os.Args) > 1 {
		x = strings.Join(os.Args[1:], " ")
	}
	if x == "" {
		x, err = internal.ReadLine(os.Stdin)
	}
	if err != nil {
		log.Print(err)
		return
	}
	title.Make(x)
}
