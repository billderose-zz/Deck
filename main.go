package main

import (
	"deck"
	"fmt"
)

func main() {
	deck := deck.NewDeck()
	deck.Shuffle()
	fmt.Println(deck)
}
