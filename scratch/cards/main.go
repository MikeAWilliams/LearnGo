package main

import "fmt"

func main() {
	cards := []string{newCard(), "Ace of Diamonds"}
	cards = append(cards, "Siz of Spades")

	for index, card := range cards {
		fmt.Println(index, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
