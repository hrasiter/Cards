package main

func main() {
	//cards := newDeck()

	//cards.saveToFile("MyCards")

	newcards := newDeckFromFile("MY")

	newcards.print()
}
