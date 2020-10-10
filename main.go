package main

func main() {
	cards := newDeck()
	// cards.saveToFile("cards.txt")
	// newDeckFromFile("cards.txt").print()
	cards.shuffle()
	cards.print()
}
