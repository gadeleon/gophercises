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
	// Make Sure Len of options is correct
	o := s.Instances["debate"].Options
	ol := len(o)
	expected := 3

	if ol != expected {
		t.Errorf("Got %d, Expected 3", ol)
	}

}
