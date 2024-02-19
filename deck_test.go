package main

import (
	"os"
	"testing"
)

func TestNewdeck(t *testing.T) {

	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeck().newDeckFromFile("_decktesting")

	if len(loadedDeck) != 20 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")

}
