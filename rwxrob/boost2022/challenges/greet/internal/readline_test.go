package internal_test

import (
	"fmt"
	"log"
	"strings"

	"github.com/mischavandenburg/go/rwxrob/boost2022/challenges/greet/internal"
)

func ExampleReadLine() {
	sr := strings.NewReader("Sample\r\n")
	line, err := internal.ReadLine(sr)
	if err != nil {
		log.Printf("error: %v", err)
		return
	}
	fmt.Printf("%q", line)
	// Output:
	// "Sample"
}
