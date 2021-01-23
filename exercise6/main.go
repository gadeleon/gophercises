package main

import (
	"encoding/json"
	"fmt"

	//"net/http"

)

//StoryJson JSON struct
type StoryJson struct {
	StoryArc string

}

type StoryArc struct {
	Title    string   `json:"title"`
	Story    []string `json:"story"`
	Options  []Option `json:"options"`

}

//Option Story arc option
type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
type Story struct {
	Instances map[string]StoryArc
}

// Parse & unmarshal json of story struct
func parseStory(jdata []byte) (Story, error) {
	a := Story{}
	err := json.Unmarshal(jdata, &a.Instances)
	if err != nil {
		// There is a better way to do this. Sorry.
		return Story{}, err
	}
	return a, nil
}
func main() {
	// PreProcess JSON
	//fmt.Println(JSONblob)
	s,err := parseStory(JSONblob)
	if err != nil {
		fmt.Println("We has error", err)
	}
	fmt.Println(s)



}
