package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	d := newDeck()

	expected := 20
	actual := len(d)

	if actual != expected {
		t.Errorf("Expected deck length of 20, but got %v", actual)
	}

	if d[0] != "Ace of Clubs" {
		t.Errorf("Expected first card to be Ace of Clubs but got %v", d[0])
	}

	if d[len(d)-1] != "King of Spades" {
		t.Errorf("Expected last card to be King of Spades but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_descktesting")

	deck := newDeck()
	deck.saveToFile("_descktesting")

	loadedDeck := newDeckFromFile("_descktesting")

	if len(loadedDeck) != 20 {
		t.Errorf("Expected deck length of 20, but got %v", len(loadedDeck))
	}

	os.Remove("_descktesting")
}
