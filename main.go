package main

func main() {
	// cards := newDeck().newDeckFromFile("first_file")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
