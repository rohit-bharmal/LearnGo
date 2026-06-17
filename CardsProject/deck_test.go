package main

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()
	if len(deck) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(deck))
	}

	if deck[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades, but got %v", deck[0])
	}

	if deck[len(deck)-1] != "King of Clubs" {
		t.Errorf("Expected King of Clubs, but got %v", deck[len(deck)-1])
	}
}
