package main

import (
	"fmt"
	"localhost/greetings"
	"log"
)

type fn func(string)(string, error)

func greetWithHandlingErrors(greet fn, name string) {
	message, err := greet(name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}

func main() {
	// get greeting message and print it
	log.SetPrefix("GREETING:")
	log.SetFlags(0)
	greetWithHandlingErrors(greetings.Hello, "Lukasz")
	greetWithHandlingErrors(greetings.RandomGreetings, "Mira")
	greetWithHandlingErrors(greetings.Hello, "")
}