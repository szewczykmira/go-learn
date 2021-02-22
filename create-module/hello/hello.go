package main

import (
	"fmt"
	"localhost/greetings"
	"log"
)

func greetWithHandlingErrors(name string) {
	message, err := greetings.Hello(name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}

func main() {
	// get greeting message and print it
	log.SetPrefix("GREETING:")
	log.SetFlags(0)
	greetWithHandlingErrors("Lukasz")
	greetWithHandlingErrors("")
}