package deck

import (
	"testing"
)


func TestNewDeckLen(t *testing.T) {
	d := New()
	b:= len(d)
	if b != 53 {
		t.Errorf("Got %d, want 53", b)
	}
}

// Test 13 of each suite
func TestNewDeckSuite(t *testing.T){
	// okay
	d := New()
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
}

// Test for a joker


