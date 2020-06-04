package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	cards := newDeck()

	fmt.Println("shuffled deck")
	shuffle := cards.shuffle(rand.New(rand.NewSource(time.Now().UnixNano())))
	shuffle.print()
}
