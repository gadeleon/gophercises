package deck

// Create a Card Type
import (
	"fmt"
)

// Note for me, J - 11, Q - 12, K - 13
type Card struct {
	suite string
	id    string
	value int
}

// What's in a CARD
// Suite
// ID
// Value

// Create a func New to make a stroebought deck o' cards.
// 52 cards + 2 joker in []card
func New() []Card {
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


		// }
		// c[s]= card{id:11}
		// c[s]= card{id:11}
		// c[s]= card{id:11}

	}
	c = append(c, Card{suite: "j", id: "0", value: 0})
	c = append(c, Card{suite: "J", id: "0", value: 0})


	return c
}
