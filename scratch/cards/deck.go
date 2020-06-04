package main

import (
	"fmt"
	"strings"
)

type deck []string

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
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

func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

//func (d dec) save(filePath string) {

//}
