package main

import (
	"net/http"
	"regexp"
	"html/template"
)
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

func renderTemplate(filename string, page *Page, w http.ResponseWriter) {
	err := templates.ExecuteTemplate(w, filename, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func getTitle(request *http.Request, w http.ResponseWriter, prefix string) (string, error) {
	return request.URL.Path[len(prefix):], nil
}
