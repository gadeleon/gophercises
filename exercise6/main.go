package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
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

type HttpHandler struct {
	Story Story
}

func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	tmpl := template.Must(template.ParseFiles("layout.html"))
	p := req.URL.Query()
	arc := p.Get("arc")
	if arc == "" {
		tmpl.Execute(res, h.Story.Instances["intro"])
	} else {
		tmpl.Execute(res, h.Story.Instances[arc])
	}

}

// Parse & unmarshal json of story struct
func parseStory(jdata []byte) (Story, error) {
	a := Story{}
	err := json.Unmarshal(jdata, &a.Instances)
	if err != nil {
		return Story{}, err
	}
	return a, nil
}
func main() {

	// PreProcess JSON
	s, err := parseStory(JSONblob)
	if err != nil {
		fmt.Println("We has error", err)
	}
	handler := HttpHandler{Story: s}
	// listen and serve
	fmt.Println("Running Server on 9777")
	http.ListenAndServe(":9777", handler)

}

// TODO: Multiples stories
// api support first choice for ken.
// Resarts / quiting / nav / startover button when finished
// breadcrumbs - adding previous & forward options.
