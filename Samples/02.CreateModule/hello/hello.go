package main

import (
	"fmt"

	greetings "example.com/greetings"
)

func main() {
	message := greetings.Hello("JB")
	fmt.Printf(message)
}
