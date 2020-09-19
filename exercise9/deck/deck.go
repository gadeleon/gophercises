package deck

// Create a Card Type
import (
	"fmt"
	"sort"
)

// Note for me, J - 11, Q - 12, K - 13
type Card struct {
	suite string
	id    string
	value int
}

type Deck struct {
	deck []Card
}

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
// FIXME: Return Deck struct
func NewStandard() []Card {
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

	return c
}

func (d Deck) Sort() {
	sort.SliceStable(d.deck, func(i, j int) bool {
		return d.deck[i].suite < d.deck[j].suite
	})
}

// TODO: Add a way to Sort by Suite (id) instead of value

// func (d *Deck) Sort() Deck {}
// TODO: "A default comparison function that can
// be used with the sorting option." (Wat?)
// TODO: Add Shuffle Method
