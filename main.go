package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Todo struct {
	Title string
	Done  bool
}

func main() {
	var addr string

	flag.StringVar(&addr, "addr", ":8000", "listen address")

	flag.Parse()

	mux := http.NewServeMux()
	// Parse all the layouts
	base := template.Must(template.ParseGlob("base.html"))

	// Clone the layouts and add the login template
	home := template.Must(template.Must(base.Clone()).ParseFiles("home.html"))
	mux.HandleFunc("GET /home", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		// Render the dashboard template with the "home" layout
		home.ExecuteTemplate(w, "base.html", map[string]string{
			"Title": "Home",
			"Name":  "Ulf",
		})
	})

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})
	fs := http.FileServer(http.Dir("static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fs))

	todos := template.Must(template.Must(base.Clone()).ParseFiles("todos.html"))
	mux.HandleFunc("GET /todos", func(w http.ResponseWriter, r *http.Request) {
		data := []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		}

		todos.ExecuteTemplate(w, "base.html", map[string]any{
			"Title": "Home",
			"Todos": data,
		})
	})
	log.Printf("Listening on %s", addr)

	http.ListenAndServe(addr, mux)
}
