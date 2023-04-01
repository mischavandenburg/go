package internal_test

import (
	"fmt"

	args "github.com/mischavandenburg/go/rwxrob/boost2022/challenges/args/internal"
)

// const testone = `
// $0  --> "./args"
// $1  --> "first"
// $2  --> "second"
// $3  --> "third"
// `

func ExampleOutput() {
	fmt.Println(args.Output("./args", "first", "second", "third"))
	// Output:
	// $0 --> "./args"
	// $1 --> "first"
	// $2 --> "second"
	// $3 --> "third"
}

/*
func TestOut(t *testing.T) {
	out := output()
	if out != "this" {
		t.Errorf("\nwant: %q\ngot:  %q\n", testone, out)

	}
}

*/
