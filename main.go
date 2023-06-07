package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

type AnotherPage struct {
	PageTitle string
	Subtitle  string
}

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
	p := os.Getenv("PORT")

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	tmpl := template.Must(template.ParseFiles("templates/layout/index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}
		tmpl.Execute(w, data)
	})
	another := template.Must(template.ParseFiles("templates/pages/another/index.html"))
	http.HandleFunc("/another", func(w http.ResponseWriter, r *http.Request) {
		data := AnotherPage{
			PageTitle: "My Another page",
			Subtitle:  "Some subtitle",
		}
		another.Execute(w, data)
	})
	fmt.Printf("The app is listening on port %v", p)
	http.ListenAndServe(p, nil)
}
