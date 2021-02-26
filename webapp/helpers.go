package main

import (
	"errors"
	"html/template"
	"net/http"
	"regexp"
)

var templates = template.Must(template.ParseFiles("tmpl/edit.html", "tmpl/view.html"))

func renderTemplate(filename string, page *Page, w http.ResponseWriter) {
	err := templates.ExecuteTemplate(w, filename, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func getTitle(request *http.Request, w http.ResponseWriter) (string, error) {
	m := validPath.FindStringSubmatch(request.URL.Path)
	if m == nil {
		http.NotFound(w, request)
		return "", errors.New("Invalid title")
	}
	return m[2], nil // The title is the second subexpression.
}
