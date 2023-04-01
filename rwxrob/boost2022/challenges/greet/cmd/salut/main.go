package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mischavandenburg/go/rwxrob/boost2022/challenges/cli"
	"github.com/mischavandenburg/go/rwxrob/boost2022/challenges/greet"
)

func main() {
	var name string
	var err error
	if len(os.Args) > 1 {
		name = strings.Join(os.Args[1:], " ")
	}
	if name == "" {
		fmt.Println("Bonjour, comment vous appelez-vous?")
		name, err = cli.ReadLine(os.Stdin)
		if err != nil {
			log.Print(err)
			return
		}
	}
	greet.Salut(name)
}
