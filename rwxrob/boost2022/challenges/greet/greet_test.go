package greet_test

import (
	"github.com/mischavandenburg/go/rwxrob/boost2022/challenges/greet"
)

func ExampleGreet_with_Arguments() {
	greet.Hi("Rob")
	// Output:
	// Hi, Rob, nice to meet you!
}
