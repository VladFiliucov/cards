package main

import (
	"testing"
)

func TestNewDeck(t *testing.T) {

	d := newDeck()
	expected := 20
	actual := len(d)

	if actual != expected {
		t.Errorf("Expected deck length of 20, but got %d", actual)
	}
}
