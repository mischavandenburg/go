package main

// import "fmt"

func main() {
	cards := newDeck()

	// deal returns 2 different values, the first will be assigned to hand, the second to remainingCards.
	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()

}
