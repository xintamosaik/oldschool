package main

import (
	"flag"

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
	fs := http.FileServer(http.Dir("static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fs))

	// Parse all the layouts
	base := template.Must(template.ParseGlob("base.html"))

	mux.HandleFunc("GET /", reactHome(base))
	mux.HandleFunc("GET /todos", reactTodos(base))

	log.Printf("Listening on %s", addr)

	http.ListenAndServe(addr, mux)
}
