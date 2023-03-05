package main

import (
	"html/template"
	"net/http"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	tmpl := template.Must(template.ParseFiles("layout.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "20230221",
			Todos: []Todo{
				{Title: "数据库", Done: false},
				{Title: "代码", Done: true},
				{Title: "配置", Done: true},
			},
		}
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":80", nil)
}
