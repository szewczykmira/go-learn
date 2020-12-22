package main

import (
	"fmt"
	"log"
	"github.com/szewczykmira/greetings"
)

func setLogger() {
	log.SetPrefix("[greetings] => ")
	log.SetFlags(0)
}

func logError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	setLogger()

	message, err := greetings.HelloFromTheOtherSide("")
	logError(err)

	fmt.Println(message)
}