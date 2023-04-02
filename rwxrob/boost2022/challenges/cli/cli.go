package cli

import (
	"bufio"
	"fmt"
	"strings"
)

// Prompt prints the optional prompt (> by default) and returns the
// string entered by the user.
func Prompt(in bufio.Reader, a string) (string, error) {
	p := "> "
	if a != "" {
		p = a
	}
	fmt.Print(p)
	return ReadLine(in)
}

// ReadLine takes any io.Reader and returns a trimmed string (initial
// and trailing white space) or an empty string and error if any error
// is encountered.
//
//	out, err := bufio.NewReader(in).ReadString('\n')
func ReadLine(in bufio.Reader) (string, error) {
	out, err := in.ReadString('\n')
	out = strings.TrimSpace(out)
	return out, err
}
