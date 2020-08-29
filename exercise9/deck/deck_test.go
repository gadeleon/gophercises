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

// Test for a joker


