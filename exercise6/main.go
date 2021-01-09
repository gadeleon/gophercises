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
type App struct {
	Instances map[string]StoryArc
}

func main() {
	// PreProcess JSON
	//fmt.Println(JSONblob)
	a := App{}
	err := json.Unmarshal(JSONblob, &a.Instances)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(a)
	}
}
