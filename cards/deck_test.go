package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected 16 but got %v", len(d))
	}

	if d[0] != "Ace Of Spades" {
		t.Errorf("Expected first card to be Ace Of Spades, but git %v", d[0])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	filename := "_decktesting"
	os.Remove(filename)

	deck := newDeck()
	deck.saveToFile(filename)
	loadedDeck := newDeckFromFile(filename)
	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards, but got %v", len(loadedDeck))
	}

	os.Remove(filename)
}
