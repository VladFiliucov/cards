package main

func main() {
	cards := newDeck()

	hand, restOfDeck := deal(cards, 2)
	hand.print()
	restOfDeck.print()
}
