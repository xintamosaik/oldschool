package main

import (
	"context"
	"github.com/a-h/templ"
	"html/template"
	"log"
	"net/http"
	 

)
var name = "Ulf Dellbruegge"

func main() {
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	home := Home()
	homeHTML, err := templ.ToGoHTML(context.Background(), home)
	if err != nil {
		log.Fatalf("failed to convert to html: %v", err)
	}
	
	index := template.Must(template.ParseFiles("index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		index.Execute(w, homeHTML)
	})

	http.HandleFunc("/cv/edit/name", func(w http.ResponseWriter, r *http.Request) {
		CVEditName(name).Render(r.Context(), w)
	})

	http.HandleFunc("/cv/update/name", func(w http.ResponseWriter, r *http.Request) {
		new := r.FormValue("name")
	
		CVShowName(new).Render(r.Context(), w)
	})

	http.ListenAndServe(":8080", nil)
}
