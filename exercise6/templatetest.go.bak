package main

import (
	//"fmt"
	"net/http"
	"html/template"
)

// THESE HAVE TO BE CAPITALIZED
type Dog struct {
	Name string
	Breed string
}
//func dog(w http.ResponseWriter, req *http.Request) {
//
//	fmt.Fprintf(w, "hello\n")
//}

func main() {
	tmpl := template.Must(template.ParseFiles("layout.html"))
	http.HandleFunc("/dog", func(w http.ResponseWriter, r *http.Request) {
		//data := "BANANAPHONE"
		//d := Dog{"Lassy", "collie"}
		//data := d
		data := []Dog{
			{Name: "Lassy", Breed: "collie"},
			{Name: "Rintintin", Breed: "mutt"},
		}
		tmpl.Execute(w, data)
	})


	http.ListenAndServe(":8091", nil)
}
//
//func main() {
//	tmpl := template.Must(template.ParseFiles("layout.html"))
//	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//		data := TodoPageData{
//			PageTitle: "My TODO list",
//			Todos: []Todo{
//				{Title: "Task 1", Done: false},
//				{Title: "Task 2", Done: true},
//				{Title: "Task 3", Done: true},
//			},
//		}
//		tmpl.Execute(w, data)
//	})
//	http.ListenAndServe(":80", nil)
//}
//
