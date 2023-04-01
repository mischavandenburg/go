package internal

import (
	"fmt"
	"strings"
)

// note this is another variadic function with the ... operator

func Output(args ...string) string {
	// fmt.Printf("%v\n%T", os.Args, os.Args)
	var buf string
	buf += strings.Join(args, " ") + "\n"
	for n, val := range args {
		buf += fmt.Sprintf("$%v --> %q\n", n, val)
	}
	return buf
}
