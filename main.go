package main

import (
	"context"
	"github.com/a-h/templ"
	"html/template"
	"log"
	"net/http"
	"time"
)

func main() {
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	templComponent := greeting()
	html, err := templ.ToGoHTML(context.Background(), templComponent)
	if err != nil {
		log.Fatalf("failed to convert to html: %v", err)
	}

	
	index := template.Must(template.ParseFiles("index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		index.Execute(w, html)
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		hello().Render(r.Context(), w)
	})
	http.HandleFunc("/clicked", func(w http.ResponseWriter, r *http.Request) {
		counts(2, 2).Render(r.Context(), w)
	})
	http.Handle("/time", templ.Handler(timeComponent(time.Now())))
	http.Handle("/404", templ.Handler(notFoundComponent(), templ.WithStatus(http.StatusNotFound)))

	http.ListenAndServe(":8080", nil)
}
