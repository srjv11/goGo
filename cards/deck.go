package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type deck []string

func deal(d deck, sizeOfHand int) (deck, deck) {
	return d[:sizeOfHand], d[sizeOfHand:]
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func newDeck() deck {

	cards := deck{}

	cardSuite := []string{"Clubs", "Diamonds", "Hearts", "Spades"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuite {
		for _, vals := range cardValues {
			cards = append(cards, vals+" of "+suit)
		}
	}

	return cards

}
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func readFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		//option1 - log the error and call newDeck()
		//option2 - log the error and entirly quit the program

		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	return strings.Split(string(bs), ",")
}
