package eightball_test

import (
	"fmt"
	"log"

	"github.com/mischavandenburg/go/rwxrob/boost2022/challenges/eightball"
)

func ExampleRespond() {
	resp := eightball.Respond()
	log.Println(resp)
	fmt.Println(len(resp) > 0)

	// Output:
	// true
}
