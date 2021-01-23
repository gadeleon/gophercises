package main

import (
	"fmt"
	"testing"
)

func TestParseStory(t *testing.T) {
	s, _ := parseStory(JSONblob)
	title := s.Instances["new-york"].Title
	if title != "Visiting New York" {
		t.Error("Got:", title, "expected:", "Visiting New York" )
	}
	// Make Sure Len of options is correct
	o := s.Instances["debate"].Options
	ol := len(o)
	fmt.Println(ol)
	if ol == 4 {
		//t.Error("Got",ol,"Expected", 4)
		t.Errorf("Got %d, Expected 4", ol)
	}
}