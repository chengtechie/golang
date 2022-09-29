package main

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Deck length is not right, current length is %v", len(d))
	}
	if d[0] != "Spades Ace" {
		t.Errorf("First card is not right, current card is %v", d[0])
	}
}
