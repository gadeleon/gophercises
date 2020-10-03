package deck

// Create a Card Type
import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
)

// Refactor for card type
type Cards struct {
	type Suit
	type Rank
}

type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

type Rank uint8

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

// Deprecated
//type Deck struct {
//	deck []Card
//}

// Make N number of decks
func New(n int) Deck {
	var d Deck

	for i := 0; i < n; i++ {
		c := NewStandard()
		d.deck = append(d.deck, c...)
	}
	return d
}

// Create a func New to make a store-bought deck o' cards.
// 52 cards + 2 joker in []card

// add a function argument of options
func NewStandard(opts ...func([]Card) []Card) []Card {

	var c []Card
	suites := []string{"H", "S", "C", "D"}
	for _, s := range suites {
		// Add Ace - 10
		for i := 0; i < 10; i++ {
			c = append(c, Card{suite: s, id: fmt.Sprint(i), value: i})
		}
		// Jack
		c = append(c, Card{suite: s, id: "11", value: 11})
		// Queen
		c = append(c, Card{suite: s, id: "12", value: 12})
		// King
		c = append(c, Card{suite: s, id: "13", value: 13})

	}
	c = append(c, Card{suite: "j", id: "0", value: 0})
	c = append(c, Card{suite: "J", id: "0", value: 0})

	for _, opt := range opts {
		c = opt(c)
	}

	return c
}

func cSort(c []Card) []Card {
	sort.SliceStable(c, func(i, j int) bool {
		return c[i].suite < c[j].suite
	})
	return c
}

func (d Deck) Sort() {
	sort.SliceStable(d.deck, func(i, j int) bool {
		return d.deck[i].suite < d.deck[j].suite
	})
}

func (d Deck) Shuffle() {
	// TODO: Figure out, where we init the seed
	// in the general case.
	rand.Shuffle(len(d.deck), func(i, j int) {
		d.deck[i], d.deck[j] = d.deck[j], d.deck[i]
	})
}

func (d *Deck) AddJokers(J, j int) error {
	// Make sure J and j are positive.
	if j < 0 || J < 0 {
		return errors.New("donuts")
	}

	for i := 0; i < j; i++ {
		d.deck = append(d.deck, Card{suite: "j", id: "0", value: 0})
	}

	for i := 0; i < J; i++ {
		d.deck = append(d.deck, Card{suite: "J", id: "0", value: 0})
	}

	return nil

}

func (d *Deck) FilterCard(c Card) error {
	// a = append(a[:i], a[i+1:]...)
	for i, j := range d.deck {
		if j == c {
			d.deck = append(d.deck[:i], d.deck[i+1:]...)
		}
	}
	return nil
}

// TODO: Add filter out specific cards from Deck.
