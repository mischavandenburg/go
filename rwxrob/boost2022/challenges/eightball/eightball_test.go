package eightball_test

import (
	"fmt"
	"log"
)

func ExampleRespond() {
	resp := eightball.Respond()
	log.Println(resp)
	fmt.Println(len(resp) > 0)
}
