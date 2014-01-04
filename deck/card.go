package deck

import "fmt"

type Suit int

type Rank int

type Card struct {
	rank  Rank
	suit  Suit
	order int
}

const (
	Ace Rank = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	Diamonds Suit = iota
	Clubs
	Hearts
	Spades
)

func (s Suit) String() string {

	switch s {
	case 0:
		return "Diamonds"
	case 1:
		return "Clubs"
	case 2:
		return "Hearts"
	case 3:
		return "Spades"
	}
	panic("Unrecognized suit")
}

func (r Rank) String() string {
	switch r {
	case 0:
		return "Ace"
	case 1:
		return "Two"
	case 2:
		return "Three"
	case 3:
		return "Four"
	case 4:
		return "Five"
	case 5:
		return "Six"
	case 6:
		return "Seven"
	case 7:
		return "Eight"
	case 8:
		return "Nine"
	case 9:
		return "Ten"
	case 10:
		return "Jack"
	case 11:
		return "Queen"
	case 12:
		return "King"
	}
	panic("Unrecognized rank")
}

func (c Card) String() string {
	return fmt.Sprintf("%s of %s\n", c.rank, c.suit)
}
