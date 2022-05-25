package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", echo)
	http.ListenAndServe(":8080", nil)
}

func echo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<html><body>")
	fmt.Fprintf(w, "<h1>Path</h1>%s", r.URL.Path)
	fmt.Fprint(w, "<h1>Cookies</h1><ul>")
	for _, c := range r.Cookies() {
		fmt.Fprintf(w, "<li>%v</li>", c)
	}
	fmt.Fprint(w, "</ul>")
	fmt.Fprint(w, "<h1>Headers</h1>")
	fmt.Fprint(w, "<ul>")
	for k, v := range r.Header {
		fmt.Fprintf(w, "<li>%s=%v</li>", k, v)
	}
	fmt.Fprint(w, "</ul>")
	fmt.Fprint(w, "</body></html>")
}
