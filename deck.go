package main

import (
	"fmt"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	allSuits := []string{"Clubs", "Diamonds", "Hearts", "Spades"}
	allValues := []string{"Ace", "Deuce", "Jack", "Queen", "King"}

	for _, suit := range allSuits {
		for _, value := range allValues {
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
