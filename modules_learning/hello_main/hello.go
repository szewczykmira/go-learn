package main

import (
	"fmt"

	"github.com/szewczykmira/greetings"
)

func main() {
	message := greetings.HelloFromTheOtherSide("Greta")
	fmt.Println(message)
}