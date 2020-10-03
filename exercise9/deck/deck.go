package deck

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

func absRank(cards Cards) int {
	return int(cards.Suit) * int(maxRank) + int(cards.Rank)
}
// Deprecated
//type Deck struct {
//	deck []Cards
//}

//// Make N number of decks
//func New(n int) Deck {
//	var d Deck
//
//	for i := 0; i < n; i++ {
//		c := NewStandard()
//		d.deck = append(d.deck, c...)
//	}
//	return d
//}

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

// suit * maxRank + Rank
//func cSort(c []Cards) []Cards {
//	sort.SliceStable(c, func(i, j int) bool {
//		return c[i].suite < c[j].suite
//	})
//	return c
//}
//
//func (d Deck) Sort() {
//	sort.SliceStable(d.deck, func(i, j int) bool {
//		return d.deck[i].suite < d.deck[j].suite
//	})
//}
//
//func (d Deck) Shuffle() {
//	// TODO: Figure out, where we init the seed
//	// in the general case.
//	rand.Shuffle(len(d.deck), func(i, j int) {
//		d.deck[i], d.deck[j] = d.deck[j], d.deck[i]
//	})
//}
//
//func (d *Deck) AddJokers(J, j int) error {
//	// Make sure J and j are positive.
//	if j < 0 || J < 0 {
//		return errors.New("donuts")
//	}
//
//	for i := 0; i < j; i++ {
//		d.deck = append(d.deck, Cards{suite: "j", id: "0", value: 0})
//	}
//
//	for i := 0; i < J; i++ {
//		d.deck = append(d.deck, Cards{suite: "J", id: "0", value: 0})
//	}
//
//	return nil
//
//}
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
