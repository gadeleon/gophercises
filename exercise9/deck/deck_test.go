package deck

import (
	"testing"
)

//// Test creation of arbitrary count of decks
func TestNew(t *testing.T) {

	for i := 0; i < 4; i++ {
		d := New(i)
		if len(d.deck) != (54 * i) {
			t.Errorf("Got %d card length, expected %d", len(d.deck), 54*1)
		}
	}
}

// Test 13 of each suite & 2 Jokers,  AKA normal deck
func TestNewDeckSuite(t *testing.T) {
	var d = NewStandard()
	c := make(map[string]int, len(d))
	for _, j := range d {
		c[j.suite]++
	}
	if c["H"] != 13 {
		t.Errorf("Got %d Hearts, want 13", c["H"])
	}
	if c["S"] != 13 {
		t.Errorf("Got %d Spades, want 13", c["S"])
	}
	if c["D"] != 13 {
		t.Errorf("Got %d Diamonds, want 13", c["D"])
	}
	if c["C"] != 13 {
		t.Errorf("Got %d Clubs, want 13", c["C"])
	}
	if c["J"] != 1 {
		t.Errorf("Got %d JOKER, want 1", c["J"])
	}
	if c["j"] != 1 {
		t.Errorf("Got %d joker, want 1", c["J"])
	}
}

// Sort Test, Just printing out.
func TestDeckSort(t *testing.T) {
	// Generate deck
	d := New(1)
	c1 := Card{"C", "0", 0}
	d.Sort()

	if d.deck[0] != c1 {
		t.Errorf("Got %v expected %v ", d.deck[0], c1)

	}

}

func TestDeckShuffle(t *testing.T) {
	d := New(1)
	t.Error(d.deck)
	d.Shuffle()
	t.Error(d.deck)

}