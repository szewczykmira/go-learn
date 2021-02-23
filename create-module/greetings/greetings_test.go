package greetings

import (
	"testing"
)


func TestCalculateLength(t *testing.T) {
	names := make(map[string]int)
	names["Lukas"] = 5
	names["Gladys"] = 6
	names[""] = 0
	for name, length := range(names){
		result := CalculateLength(name)
		if result != length{
			t.Fatalf("Name: %v. Expected: %d got: %d", name, length, result)
		}
	}
}


func TestHelloWhenNamePassed(t *testing.T) {
	names := make(map[string]string)
	names["Lukas"] = "Hi, Lukas. Welcome! Do you know that your name has 5 letters?"
	names["Gladys"] = "Hi, Gladys. Welcome! Do you know that your name has 6 letters?"
	for name, response := range(names){
		result, _ := Hello(name)
		if result != response{
			t.Fatalf("Name %v. Got %v", name, result)
		}
	}
}

func TestHelloWhenEmptyStringPassed(t *testing.T) {
	name := ""
	message, err := Hello(name)
	if err == nil {
		t.Fatalf("Error not raised, got: %v", message)
	}
}