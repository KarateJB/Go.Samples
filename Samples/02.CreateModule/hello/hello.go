package main

import (
	"fmt"
	"log"

	greetings "example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(log.LstdFlags | log.Lshortfile) // Set to 1 as default format, or use constant flags, see https://golang.org/pkg/log/#pkg-constants

	message, err := greetings.Hello("JB")

	if err != nil {
		log.Fatal(err)
		// Output "greetings: 2021/01/20 11:04:10 hello.go:17: Empty name"
	}

	fmt.Printf(message)
}
