package eightball

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/mischavandenburg/go/rwxrob/boost2022/challenges/cli"
)

var Responses = [...]string{
	"Yes",
	"No",
	"Maybe",
	"Why not?",
}

func Respond() string {
	rand.Seed(time.Now().UnixNano())
	return Responses[rand.Intn(len(Responses))]
}

func Run() {
	fmt.Println("Welcome to the magic eight ball!\nwhat is your question?")
	for {
		cli.Prompt(os.Stdin, "🎱 ")
		fmt.Println(Respond())
	}
}
