package main

import (
	"math/rand"
	"sort"
	"time"
)

type Deck [52]Card

func NewDeck() *Deck {
	var deck Deck
	for i := 0; i < 52; i++ {
		rand.Seed(time.Now().UTC().UnixNano())
		deck[i] = Card{rank: Rank(i % 13), suit: Suit(i % 4), order: rand.Intn(52)}
	}
	return &deck
}

func (d *Deck) Len() int {
	return len(d)
}

func (d *Deck) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (d *Deck) Less(i, j int) bool {
	return d[i].order < d[j].order
}

func (d *Deck) Shuffle() {
	sort.Sort(d)
}

func (d *Deck) String() string {
	s := ""
	for _, card := range d {
		s += card.String()
	}
	return s
}
