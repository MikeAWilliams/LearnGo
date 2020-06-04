package main

import "testing"

func uniqueCards(input deck) deck {
	foundAlready := map[string]bool{}
	result := deck{}
	for _, card := range input {
		if !foundAlready[card] {
			result = append(result, card)
			foundAlready[card] = true
		}
	}
	return result
}
func TestNewDeck(t *testing.T) {
	testObject := newDeck()
	if 52 != len(testObject) {
		t.Errorf("Expected len(testObject) of 52 but got %d", len(testObject))
	}

	uniqueCardsInTestObject := uniqueCards(testObject)
	if 52 != len(uniqueCardsInTestObject) {
		t.Errorf("Expected len(uniqueCardsInTestObject) of 52 but got %d", len(uniqueCardsInTestObject))
	}

}
