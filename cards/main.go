package main

import "fmt"

func main() {
	card := newCard()
	fmt.Println(card)
}

// the word string tells the compiler that the function will always return a string
func newCard() string {
	return "Five of Diamonds"
}
