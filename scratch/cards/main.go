package main

func main() {
	cards := deck{newCard(), "Ace of Diamonds"}
	cards = append(cards, "Siz of Spades")
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
