package main

import (
	"fmt"
)

func main() {
	cards := newDeck()
	fmt.Println("cards")
	cards.print()

	fmt.Println("")
	fmt.Println("shuffled deck")
	shuffle := cards.shuffle()
	shuffle.print()
}
