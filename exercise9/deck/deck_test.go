package deck

import (
	// "fmt"
	"testing"
)

// EqualCards compares slices for us.
// TY: https://yourbasic.org/golang/compare-slices/
func EqualCards(a, b []Cards) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if absRank(v) != absRank(b[i]) {
			return false
		}
	}
	return true
}

// Test creation of arbitrary count of decks
// Repurposed from old New!!
func TestNew(t *testing.T) {
	d := New()
	if len(d) != 52 {
		t.Errorf("Got %d cards, want %s", len(d), "52")
	}

}

func TestAbsRank(t *testing.T) {
	c1 := Cards{Spade, Jack}
	c2 := Cards{Suit: Diamond, Rank: Six}
	c3 := Cards{Suit: Club, Rank: Ace}
	c4 := Cards{Suit: Heart, Rank: Four}
	var vals = [...]Cards{c1, c2, c3, c4}
	for _, card := range vals {
		abs := absRank(card)
		exp := int(card.Suit)*int(maxRank) + int(card.Rank)
		if abs != exp {
			t.Errorf("Got %d for %v, want %d", abs, card, exp)
		}
	}

}

func TestShuffle(t *testing.T) {
	// Semi-deterministic seed
	d := New(Shuffle)
	e := New()
	result := EqualCards(d, e)
	if result != false {
		t.Errorf("Got %v, want false", result)
	}
}

//
//// Test 13 of each suite & 2 Jokers,  AKA normal deck
//func TestNewDeckSuite(t *testing.T) {
//	var d = New()
//	c := make(map[string]int, len(d))
//	for _, j := range d {
//		c[j.suite]++
//	}
//	if c["H"] != 13 {
//		t.Errorf("Got %d Hearts, want 13", c["H"])
//	}
//	if c["S"] != 13 {
//		t.Errorf("Got %d Spades, want 13", c["S"])
//	}
//	if c["D"] != 13 {
//		t.Errorf("Got %d Diamonds, want 13", c["D"])
//	}
//	if c["C"] != 13 {
//		t.Errorf("Got %d Clubs, want 13", c["C"])
//	}
//	if c["J"] != 1 {
//		t.Errorf("Got %d JOKER, want 1", c["J"])
//	}
//	if c["j"] != 1 {
//		t.Errorf("Got %d joker, want 1", c["J"])
//	}
//}
//
//// Sort Test, Just printing out.
//func TestDeckSort(t *testing.T) {
//	// Generate deck
//	d := New(1)
//	c1 := Card{"C", "0", 0}
//	d.Sort()
//
//	if d.deck[0] != c1 {
//		t.Errorf("Got %v expected %v ", d.deck[0], c1)
//
//	}
//
//}
//
//func TestfoDeckSort(t *testing.T) {
//	// Generate deck
//
//	d := NewStandard(cSort)
//	c1 := Card{"C", "0", 0}
//
//	if d[0] != c1 {
//		t.Errorf("Got %v expected %v ", d[0], c1)
//
//	}
//
//}
//
//func TestDeckShuffle(t *testing.T) {
//	// We want deterministic testing for tests.
//	// Explicit declaration of seed 1 even though default
//	rand.Seed(1)
//	// Generate 1 deck
//	d := New(1)
//	// Shuffle
//	d.Shuffle()
//	exp := Card{"S", "7", 7}
//	if d.deck[0] != exp {
//		t.Errorf("Got %v, expected %v", d.deck[0], exp)
//	}
//
//}
//
//func TestNewDeckJokers(t *testing.T) {
//	d := New(1)
//	err := d.AddJokers(1, 1)
//
//	// Sort the deck into suites
//	c := make(map[string]int, len(d.deck))
//	for _, j := range d.deck {
//		c[j.suite]++
//	}
//
//	if c["J"] != 2 {
//		t.Errorf("Got %d JOKER, want 2", c["J"])
//	}
//	if c["j"] != 2 {
//		t.Errorf("Got %d joker, want 2", c["J"])
//	}
//
//	// Check for negative # of jokers.
//	err = d.AddJokers(-1, 1)
//	if err != nil {
//		return
//	} else {
//		t.Errorf("Expected  got %v", err)
//	}
//
//}
//func TestFilterCard(t *testing.T) {
//	d := New(2)
//	oglen := len(d.deck)
//	// Card we want to filter out
//	filtered := Card{"H", "2", 2}
//	d.FilterCard(filtered)
//
//	for _, j := range d.deck {
//		if j == filtered {
//			t.Errorf("%v found, should not be in slice", filtered)
//		}
//	}
//	// Check to see all other expected cards are present
//	total := len(d.deck)
//	if total == oglen {
//		t.Errorf("Original length is %v, should be less than %v", oglen, total)
//	}
//
//}
