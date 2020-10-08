package main

import (
	"fmt"
)

func main() {
	cards := newDeck()

	hand, restOfDeck := deal(cards, 2)
	fmt.Println(hand.toString())
	fmt.Println(restOfDeck.toString())
	// hand.print()
	// restOfDeck.print()
}
