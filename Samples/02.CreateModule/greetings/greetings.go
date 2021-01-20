package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a welcome message for the named person.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Empty name")
	}

	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
