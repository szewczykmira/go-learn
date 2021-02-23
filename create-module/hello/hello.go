package main

import (
	"fmt"
	"localhost/greetings"
	"log"
)

type fn func(string)(string, error)

func handleErrorOrReturnValue(message string, err error){
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}

func greetWithHandlingErrors(greet fn, name string) {
	handleErrorOrReturnValue(greet(name))
}

func main() {
	// get greeting message and print it
	log.SetPrefix("GREETING:")
	log.SetFlags(0)
	names := []string{"Gladys", "Samantha", "John"}
	message, err := greetings.Hellos(names, greetings.Hello)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	greetWithHandlingErrors(greetings.Hello, "Lukasz")
	greetWithHandlingErrors(greetings.RandomGreetings, "Mira")
	greetWithHandlingErrors(greetings.Hello, "")
}