package main

import "testing"

func TestNewDeck(t *testing.T) {
	testObject := newDeck()
	if 52 != len(testObject) {
		t.Errorf("Expected length of 52 but got %d", len(testObject))
	}
}
