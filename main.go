package main

import "fmt"

func main() {
	deck := NewDeck()
	deck.Shuffle()
	fmt.Println(deck)
}
