package main

import (
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	http.Handle("/hello", templ.Handler(hello()))

	

	http.Handle("/", templ.Handler(timeComponent(time.Now())))
	http.Handle("/404", templ.Handler(notFoundComponent(), templ.WithStatus(http.StatusNotFound)))

	http.ListenAndServe(":8080", nil)
}

