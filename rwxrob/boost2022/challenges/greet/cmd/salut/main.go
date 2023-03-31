package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mischavandenburg/go/rwxrob/boost2022/challenges/greet"
	"github.com/mischavandenburg/go/rwxrob/boost2022/challenges/greet/internal"
)

func main() {
	var name string
	var err error
	if len(os.Args) > 1 {
		name = strings.Join(os.Args[1:], " ")
	}
	if name == "" {
		fmt.Println("Bonjour, comment vous appelez-vous?")
		name, err = internal.ReadLine(os.Stdin)
		if err != nil {
			log.Print(err)
			return
		}
	}
	greet.Salut(name)
}