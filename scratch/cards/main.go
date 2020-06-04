package main

import "fmt"

func main() {
	cards := newDeck()
	fmt.Println("cards")
	cards.print()

	hand, deck := deal(cards, 5)
	fmt.Println("")
	fmt.Println("hand")
	hand.print()

	fmt.Println("")
	fmt.Println("deck")
	deck.print()

	fmt.Println(deck.toString())

	deck.saveToFile("stuff.txt")

	deckFromFile := newDeckFromFile("stuff.txt")
	fmt.Println("")
	fmt.Println("loaded deck")
	deckFromFile.print()
}
