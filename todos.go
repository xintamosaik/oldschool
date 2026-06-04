package main

import (
	"html/template"
	"net/http"
)

func reactTodos(base *template.Template) func(http.ResponseWriter, *http.Request) {
	html := `
	{{define "yield"}}
		<ul>
			{{range .Todos}}
				{{if .Done}}
					<li class="done">{{.Title}}</li>
				{{else}}
					<li>{{.Title}}</li>
				{{end}}
			{{end}}
		</ul>
	{{end}}`

	todos := template.Must(template.Must(base.Clone()).Parse(html))
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		data := []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		}

		todos.ExecuteTemplate(w, "base.html", map[string]any{
			"Title": "Home",
			"Todos": data,
		})
	}

}
