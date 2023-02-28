package main

// import "fmt"

func main() {
	cards := newDeck()

	// deal returns 2 different values
	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()

}
