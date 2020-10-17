package deck

import (
	"math/rand"
	"sort"
	"time"
)

// Refactor for card type
type Cards struct {
	Suit
	Rank
}

type Suit uint8

// We know Suits never change.
// They can be constants
const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

type Rank uint8

// We know ranks can never change
// They can be constants
const (
	_ Rank = iota
	Ace
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

// We want to know the proper range of ranks
const (
	minRank = Ace
	maxRank = King
)

// Take number values of cards and create math for sorting
func absRank(cards Cards) int {
	return int(cards.Suit)*int(maxRank) + int(cards.Rank)
}

// Create a func New to make a store-bought deck o' cards.
// 52 cards + 2 joker in []Cards

// Need range of suits to iterate for card declartion.
// Think global
var suits = [...]Suit{Spade, Diamond, Club, Heart}

// add a function argument of options
func New(opts ...func([]Cards) []Cards) []Cards {
	var c []Cards
	for _, s := range suits {
		// Add Ace -
		for rank := minRank; rank <= maxRank; rank++ {
			c = append(c, Cards{Suit: s, Rank: rank})
		}
	}

	for _, opt := range opts {
		c = opt(c)
	}

	return c
}

// Shuffle Cards
// TY: https://stackoverflow.com/questions/12321133/how-to-properly-seed-random-number-generator?lastactivity
func Shuffle(c []Cards) []Cards {
	// Use time to generate random seed for each shuffle
	rand.Seed(time.Now().UTC().UnixNano())
	rand.Shuffle(len(c), func(i, j int) {
		c[i], c[j] = c[j], c[i]
	})

	return c
}

// suit * maxRank + Rank
//func cSort(c []Cards) []Cards {
//	sort.SliceStable(c, func(i, j int) bool {
//		return c[i].suite < c[j].suite
//	})
//	return c
//}
//
func Sort(c []Cards) []Cards {
	sort.SliceStable(c, func(i, j int) bool {
		return absRank(c[i]) < absRank(c[j])
	})
	return c
}

func AddJokers(n int) func([]Cards) []Cards {
	return func(c []Cards) []Cards {
		// Here we add the actual adding of cards
		for i := 0; i <= n; i++ {
			c = append(c, Cards{Suit: 4, Rank: 0})
		}
		return c

	}
}

//
//func (d *Deck) FilterCard(c Card) error {
//	// a = append(a[:i], a[i+1:]...)
//	for i, j := range d.deck {
//		if j == c {
//			d.deck = append(d.deck[:i], d.deck[i+1:]...)
//		}
//	}
//	return nil
//}
//
//// TODO: Add filter out specific cards from Deck.
