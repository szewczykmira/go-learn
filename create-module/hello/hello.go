package main

import (
	"fmt"
	"localhost/greetings"
)

func main() {
	// get greeting message and print it
	message := greetings.Hello("Misio")
	fmt.Println(message)
}