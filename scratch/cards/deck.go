package main

import (
	"fmt"
	"io/ioutil"
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

func (d deck) saveToFile(filePath string) error {
	return ioutil.WriteFile(filePath, []byte(d.toString()), 0666)
}

func newDeckFromFile(filePath string) deck {
	bytes, err := ioutil.ReadFile(filePath)
	if nil != err {
		return deck{}
	}
	stringArray := strings.Split(string(bytes), ",")
	return deck(stringArray)
}
