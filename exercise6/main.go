package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)
	//"os")



//StoryJson JSON struct
type StoryJson struct {
	StoryArc StoryArc
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

type Index struct {
	Stories []byte
}

func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}


//TODO: Add path support.
func getFileNameList(p string) ([]string, error) {
	//p,err  := os.Getwd()
	//if err != nil {
	//	log.Fatalf("Unable to list files at %s",p)
	//}
	files, err := ioutil.ReadDir(p)
	if err != nil {
		log.Fatal(err)
		return []string{}, err
	}
	flist := []string{}

	for _, f := range files {
		// TODO: Dangerous but we'll fix later (the period)
		if strings.Contains(f.Name(), ".json") {
			target := filepath.Join(p,f.Name())
			flist = append(flist, target)

		}
	}
	return flist, nil
}

func ingestFile(fn string) StoryJson {
	f, err := os.Open(fn)
	if err != nil {
		log.Panic("Oh shit")
	}
	data := make([]byte,0)

	//watch for this no :
	_, err = f.Read(data)
	story := StoryJson{}
	err = json.Unmarshal(data,story)
	return story





	//return StoryJson{}

}
//func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
//	tmpl := template.Must(template.ParseFiles("layout.html"))
//	p := req.URL.Query()
//	arc := p.Get("arc")
//	if arc == "" {
//		tmpl.Execute(res, h.Story.Instances["intro"])
//	} else {
//		tmpl.Execute(res, h.Story.Instances[arc])
//	}
//
//}
//// Parse multiple stories accepts a story
//func parseStories(jdata []byte) (Index, error) {
//	return Index{}, nil
//}
//
//// Parse & unmarshal json of story struct
//func parseStory(jdata []byte) (Story, error) {
//	a := Story{}
//	err := json.Unmarshal(jdata, &a.Instances)
//	if err != nil {
//		return Story{}, err
//	}
//	return a, nil
//}
//func main() {
//
//	// PreProcess JSON
//	s, err := parseStory(JSONblob)
//	if err != nil {
//		fmt.Println("We has error", err)
//	}
//	handler := HttpHandler{Story: s}
//	// listen and serve
//	fmt.Println("Running Server on 9777")
//	http.ListenAndServe(":9777", handler)
//
//}
//
//// api support first choice for ken.
//// Resarts / quiting / nav / startover button when finished
//// breadcrumbs - adding previous & forward options.
//// TODO: Make func to Read the files from getFileNameList()
//// TODO: "Read" or next func after read tosses into Story Struct
//// build list function, something something composition.