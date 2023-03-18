package main

import "testing"

// tests are automatically called by the go test runner with an argument of t *testing.T
// t is the test handler
func TestNewDeck(t *testing.T) {
	d := newDeck()

	// we check if the deck has 16 cards inside of it. If it doesn't,
	// we notify the test handler by calling its Errorf function
	if len(d) != 16 {
		// the %v takes the value of the variable after the string
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}
}
