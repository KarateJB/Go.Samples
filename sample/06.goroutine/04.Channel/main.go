package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int, 2)
	go write(ch)

	time.Sleep(2 * time.Second)

	// Method 1. range
	// for v := range ch {
	// 	fmt.Println("Read value", v, "from channel")
	// 	time.Sleep(2 * time.Second)
	// }

	// Method 2. for
	// LOOP1:
	// 	for {

	// 		if v, ok := <-ch; ok {
	// 			fmt.Println("Read value", v, "from channel")
	// 			time.Sleep(2 * time.Second)
	// 		} else {
	// 			break LOOP1
	// 		}
	// 	}

	// Method 3. for + select
LOOP2:
	for {
		select {
		case v, ok := <-ch:
			if ok {
				fmt.Println("Read value", v, "from channel")
			} else {
				break LOOP2
			}
		case <-time.After(1 * time.Second):
			fmt.Println("Timout...")
			// default:
			// 	fmt.Println("Channel blocking...") // This won't happen cus we break the loop when the channel is closed.
		}

		time.Sleep(2 * time.Second)
	}
}

func write(ch chan int) {

	for i := 1; i <= 6; i++ {
		time.Sleep(500 * time.Millisecond)
		ch <- i
		fmt.Println("Successfully wrote", i, "to channel")
	}

	close(ch) // Comment out this line to test the default or timeout cases in "3. for + select"
}
