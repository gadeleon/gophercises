package main

import (
	"fmt"
	"net/http"
)

//StoryJson JSON struct
type StoryJson struct {
	StoryArc string   `json:"arc"`
	Title    string   `json:"title"`
	Story    []string `json:"story"`
	Options  []Option `json:"options"`
}

//Option Story arc option
type Option struct {
	Text string `json:"arcText"`
	Arc  string `json:"nextArc"`
}

func main() {
	// PreProcess JSON

}
