package main

import (
	// "fmt"
	"testing"
)

func TestParseStory(t *testing.T) {
	s, _ := parseStory(JSONblob)
	title := s.Instances["new-york"].Title
	if title != "Visiting New York" {
		t.Error("Got:", title, "expected:", "Visiting New York")
	}
	o := s.Instances["debate"].Options
	ol := len(o)
	// Debate has 3 options
	expected := 3

	if ol != expected {
		t.Errorf("Got %d, Expected %d", ol, expected)
	}

}
