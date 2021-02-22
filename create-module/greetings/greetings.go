package greetings

import "fmt"

// CalculateLength Return length of the name
func CalculateLength(name string) int {
	return len(name)
}


// Hello will return name and length of given name
func Hello(name string) string {
	value := CalculateLength(name)
	fmt.Println(value)
	message := fmt.Sprintf("Hi, %v. Welcome! Do you know that your name has %d letters?", name, value)
	return message
}