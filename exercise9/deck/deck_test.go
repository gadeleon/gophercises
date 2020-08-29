package deck

import (
	"testing"
)

// Test for New
func TestNewDeckLen(t *testing.T) {
	d := New(2)
	b:= len(d)
	if b != 53 {
		t.Error(("ham"))
	}
}

// Length of []Card is 53
// 13 of each Suite in the Deck
// 1 Joker in the Deck
