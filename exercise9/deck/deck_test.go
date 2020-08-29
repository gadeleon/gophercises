package deck

import (
	"fmt"
	"testing"
)

// Test for New
func TestNewDeckLen(t *testing.T) {
	d := new()
	len(d) == 53
}

// Length of []Card is 53
// 13 of each Suite in the Deck
// 1 Joker in the Deck
