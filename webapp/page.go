package main

import (
	"fmt"
	"io/ioutil"
)

// Page keeps info about... pages?
type Page struct {
	Title string
	Body  []byte
}

// *Page is a reciever
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	// &Page - pointer
	return &Page{Title: title, Body: body}, nil
}

func readAndSave() {
	p1 := &Page{Title: "Test page", Body: []byte("This is sample page")}
	p1.save()
	p2, _ := loadPage("Test page")
	fmt.Println(string(p2.Body))
}
