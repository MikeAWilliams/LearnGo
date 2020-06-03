package main

import "fmt"

type deck []string

func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}

func newDeck() deck {
	suits := []string{"Spades", "Clubs", "Hearts", "Diamonds"}
	values := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	result := deck{}

	for _, suit := range suits {
		for _, value := range values {
			result = append(result, value+" of "+suit)
		}
	}
	return result
}
