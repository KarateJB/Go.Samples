package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a welcome message for the named person.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Empty name")
	}

	// message := fmt.Sprintf("Hi, %v. Welcome!", name)
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// init sets the seed of rand
func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	// A slice of message formats
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
