package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mischavandenburg/go/rwxrob/boost2022/challenges/cli"
	"github.com/mischavandenburg/go/rwxrob/boost2022/challenges/greet"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var name string
	var err error
	if len(os.Args) > 1 {
		name = strings.Join(os.Args[1:], " ")
	}
	if name == "" {
		fmt.Println("Hello there, what's your name?")
		name, err = cli.ReadLine(in)
		if err != nil {
			log.Print(err)
			return
		}
	}
	greet.Hi(name)
}
