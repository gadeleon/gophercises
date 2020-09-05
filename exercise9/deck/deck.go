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
// 52 cards + 1 joker in []card
func New() []Card {
	var c []Card
	suites := []string{"H", "S", "C", "D"}
	//loop 13, one for each suite
	for _, s := range suites {
		// Add numbered cards
		for i := 0; i > 10; i++ {
			c = append(c, Card{suite: s, id: fmt.Sprint(i), value: i})
			println(c)
		}
		// c[s{suite
		// 	id i
		// 	value i }],k,q,

		// }
		// c[s]= card{id:11}
		// c[s]= card{id:11}
		// c[s]= card{id:11}

	}

	return c
}
