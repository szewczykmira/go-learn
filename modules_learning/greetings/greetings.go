package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

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

func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	formats := []string{
		"Hi! %v Welcome!",
		"Hello %v. Long time no see",
		"Siema %v!",
	}
	return formats[rand.Intn(len(formats))]
}

// GetRandomGreeting returns some nice greetings
func GetRandomGreeting(name string) (string, error){
	if name == "" {
		// If no name was passed return error
		return "", errors.New("empty name")
	}
	var message string
	message = fmt.Sprintf(randomFormat(), name)
	return message, nil
}
