package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"html/template"

)

//StoryJson JSON struct
type StoryJson struct {
	StoryArc string
}

type StoryArc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

//Option Story arc option
type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
type Story struct {
	Instances map[string]StoryArc
}

type HttpHandler struct{
	Story Story

}


func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	tmpl := template.Must(template.ParseFiles("layout.html"))
	//data := []byte("Hello World!") // slice of bytes
	fmt.Println("Handler called!")
	// Start at intro
	tmpl.Execute(res, h.Story.Instances["intro"])

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
	s, err := parseStory(JSONblob)
	if err != nil {
		fmt.Println("We has error", err)
	}
	//fmt.Println(s)
	// Establish template
	// create a new handler
	handler := HttpHandler{Story: s}
	// listen and serve
	fmt.Println("Running Server on 9777")
	http.ListenAndServe(":9777", handler)



}
// TODO: If there's no param, return intro, else, return storyarc
// TODO: Rendor out arcs, turn them into links with arc as the key