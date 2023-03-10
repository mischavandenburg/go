package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {

		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {

	for i, card := range d {
		fmt.Println(i, card)
	}
}

// a function called deal that takes a type of deck, referenced as d, and an int of handSize. Then we tell Go that we are expecting to return two separate values of type deck.
// Go has support for multiple return values from a function. The second set of parentheses
// defines that it will return two values of type deck.
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// uses the Join function from the imported strings package to create one long string from a slice of strings
func (d deck) toString() string {
	return strings.Join([]string(d), ",")

}

// uses ioutil to write to disk. ioutil takes the filename, a byte slace and the file permissions as arguments.
// we use the toString() function to convert our deck to one long string, and then call the type conversion to byte slice.
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	// ReadFile returns a byte slice and possibly an error, so we save both into new variables
	//
	bs, err := ioutil.ReadFile(filename)

	// this code is only executed if there is an error
	// very common to see this with functions that return an error object.
	// We print out the error object and quit the program
	if err != nil {
		fmt.Println("Error:", err)
		// os.Exit causes the program to exit with the given status code. Zero = success
		os.Exit(1)
	}

	// convert the byte slice to a slice of strings and save into s
	s := strings.Split(string(bs), ",")
	// convert the slice of strings to type deck and return
	return deck(s)

}
