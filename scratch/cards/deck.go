package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
)

type deck []string

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func newDeck() deck {
	suits := [4]string{"Spades", "Clubs", "Hearts", "Diamonds"}
	values := [13]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

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

func newDeckFromFile(filePath string) (deck, error) {
	bytes, err := ioutil.ReadFile(filePath)
	if nil != err {
		return deck{}, err
	}
	stringArray := strings.Split(string(bytes), ",")
	return deck(stringArray), nil
}

func (d deck) shuffle(randomGen *rand.Rand) deck {
	result := make(deck, len(d))
	copy(result, d)

	for index := range result {
		newIndex := randomGen.Intn(len(result))
		result[newIndex], result[index] = result[index], result[newIndex]
	}
	return result
}
