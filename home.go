package main

import (
	"html/template"
	"net/http"
)

func reactHome(base *template.Template) func(http.ResponseWriter, *http.Request) {
	html := `
	{{define "yield"}}
  		<p>Welcome back, {{.Name}}!</p>
	{{end}}`
	// Clone the layouts and add the login template
	home := template.Must(template.Must(base.Clone()).Parse(html))
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		// Render the dashboard template with the "home" layout
		home.ExecuteTemplate(w, "base.html", map[string]string{
			"Title": "Home",
			"Name":  "Ulf",
		})
	}

}
