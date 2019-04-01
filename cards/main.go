package main

import "fmt"

func main() {

	cards := newCards()
	fmt.Println(cards)

}

func newCards() string {
	return "Five of Diamonds"
}
