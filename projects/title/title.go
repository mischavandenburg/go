// Package title converts a string to title case.

package title

import (
	"fmt"
	"strings"
)

func Make(s string) {
	t := strings.Title(s)
	fmt.Printf("%v\n", t)
}
