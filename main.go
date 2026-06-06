package main

import (
	"context"
	"github.com/a-h/templ"
	"html/template"
	"log"
	"net/http"

)

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

	http.HandleFunc("/todo/new", func(w http.ResponseWriter, r *http.Request) {
		TodoNew().Render(r.Context(), w)
	})
	http.ListenAndServe(":8080", nil)
}
