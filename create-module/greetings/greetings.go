package greetings

import (
	"errors"
	"fmt"
)

// CalculateLength Return length of the name
func CalculateLength(name string) int {
	return len(name)
}


// Hello will return name and length of given name
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Empty name was passed")
	}

	value := CalculateLength(name)
	message := fmt.Sprintf("Hi, %v. Welcome! Do you know that your name has %d letters?", name, value)
	return message, nil
}