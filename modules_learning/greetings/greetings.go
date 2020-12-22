package greetings

import "fmt"

// Hello gets parameter and ignores it
func Hello(name string) string {
	message := fmt.Sprintf("Hello majfriend")
	return message
}

// HelloFromTheOtherSide is original function from tutorial
func HelloFromTheOtherSide(name string) string {
	var message string
    message = fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}
