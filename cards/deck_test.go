package main

import "testing"

//import "fmt"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected 20, but got %v", len(d))
	}
}
