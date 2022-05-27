package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int, 2)
	go write(ch)

	time.Sleep(2 * time.Second)

	for v := range ch {
		fmt.Println("Read value", v, "from channel")
		time.Sleep(2 * time.Second)
	}

	// Or iterate channel like this
	// LOOP1:
	// 	for {

	// 		if v, ok := <-ch; ok {
	// 			fmt.Println("Read value", v, "from channel")
	// 			time.Sleep(2 * time.Second)
	// 		} else {
	// 			break LOOP1
	// 		}
	// 	}

	// LOOP2:
	// 	for {
	// 		select {
	// 		case v, ok := <-ch:
	// 			if ok {
	// 				fmt.Println("Read value", v, "from channel")
	// 				time.Sleep(2 * time.Second)
	// 			} else {
	// 				break LOOP2
	// 			}
	// 		default:
	// 			fmt.Println("Channel blocking...") // This won't happen cus we break the loop when the channel is closed.
	// 		}
	// 	}
}

func write(ch chan int) {

	for i := 1; i <= 6; i++ {
		time.Sleep(2 * time.Second)
		ch <- i
		fmt.Println("Successfully wrote", i, "to channel")
	}

	close(ch)
}
