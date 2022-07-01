package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type NewAggPage struct {
	Title string
	News  string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> Go is Amazing!</h1>")
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	p := NewAggPage{Title: "Amazing News Aggregator", News: "some news"}
	q := NewAggPage{Title: "I am using GPS", News: "GPS news"}
	t, _ := template.ParseFiles("basictemplating.html")
	t.Execute(w, p)
	t.Execute(w, q)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about/", newsAggHandler)
	http.ListenAndServe("8000", nil)
}
