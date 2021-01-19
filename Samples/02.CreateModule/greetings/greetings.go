package greeting

import "fmt"

// Return a welcome message for the named person.
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
