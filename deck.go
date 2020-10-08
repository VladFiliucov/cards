package main

import (
	"fmt"
	"io/ioutil"
	"strings"
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

func deal(d deck, cardsToTake int) (deck, deck) {
	return d[:cardsToTake], d[cardsToTake:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
