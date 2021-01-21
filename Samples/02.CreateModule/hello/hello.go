/*
  See [Tutorial: Create a Go module]
  https://golang.org/doc/tutorial/create-module
*/
package main

import (
	"fmt"
	"log"

	greetings "example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(log.LstdFlags | log.Lshortfile) // Set to 1 as default format, or use constant flags, see https://golang.org/pkg/log/#pkg-constants

	// Hello message for a single person
	message, err := greetings.Hello("JB")

	if err != nil {
		log.Fatal(err)
		// Output "greetings: 2021/01/20 11:04:10 hello.go:17: Empty name"
	}

	fmt.Printf(message + "\n")

	// Hello messages for a group
	names := []string{"Dog", "Cat", "Rabbit"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
		// Output "greetings: 2021/01/20 11:04:10 hello.go:17: Empty name"
	}

	for _, msg := range messages {
		fmt.Println(msg)
	}
}
