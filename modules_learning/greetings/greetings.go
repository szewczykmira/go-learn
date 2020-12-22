package greetings

import "fmt"
import "errors"

// Hello gets parameter and ignores it
func Hello(name string) string {
	message := fmt.Sprintf("Hello majfriend")
	return message
}

// HelloFromTheOtherSide is original function from tutorial
func HelloFromTheOtherSide(name string) (string, error) {
	if name == "" {
		// If no name was passed return error
		return "", errors.New("empty name")
	}
	var message string
    message = fmt.Sprintf("Hi, %v. Welcome!", name)
    return message, nil
}
