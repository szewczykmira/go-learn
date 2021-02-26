package main

import (
	"log"
	"net/http"
)

const Port = ":8088"


func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(Port, nil))
}
