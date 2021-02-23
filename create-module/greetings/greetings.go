package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
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

// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one value from the set of greetings
func randomFormat() string {
	formats := []string{
		"Hi! %v welcome",
		"Alloha %v",
		"Dzien dobry %v!",
	}
	return formats[rand.Intn(len(formats))]
}

// RandomGreetings will return random string
func RandomGreetings(name string) (string, error) {
	if name == "" {
		return "", errors.New("Empty name was passed")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos will return greetings for every name in a list.
func Hellos(names []string, greet func(string) (string, error)) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := greet(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}
