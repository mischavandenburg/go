package main

import (
	"log"
	"os"

	"github.com/mischavandenburg/go/projects/title"
	"github.com/mischavandenburg/go/projects/title/internal"
)

func main() {
	var x string
	var err error
	x, err = internal.ReadLine(os.Stdin)
	if err != nil {
		log.Print(err)
		return
	}
	title.Make(x)
}
