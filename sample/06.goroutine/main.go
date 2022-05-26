package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// NG
	// ngSample()

	// WaitGroup
	// waitGroupSample()

	// Channel
	channelSample()
}

// Output: show message
func output(s string, delay int) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Duration(delay) * time.Millisecond)
		fmt.Println(s)
	}
}

// OutputByWg: show message but use WaitGroup to wait target goroutines
func outputByWg(s string, delay int, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done() // decrease counter by 1, once counter eauals 0, WaitGroup stop blocking.
	}

	for i := 0; i < 5; i++ {
		time.Sleep(time.Duration(delay) * time.Millisecond)
		fmt.Println(s)
	}
}

// outputByChannel: show message but use channel to block
func outputByChannel(s string, delay int, ch chan string) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Duration(delay) * time.Millisecond)
		fmt.Println(s)
	}

	ch <- "Done"
}

func ngSample() {
	// The goroutines will be terminated once the main goroutine ends.
	go output("goroutine 1", 100)
	go output("goroutine 2", 300)
	output("goroutine main", 200)
}

func waitGroupSample() {
	// Wait goroutine 1 & 2.
	wg := new(sync.WaitGroup)
	wg.Add(2) // Set thread counter = 2
	go outputByWg("goroutine 1", 100, wg)
	go outputByWg("goroutine 2", 300, wg)

	outputByWg("goroutine main", 200, nil) // The process will still wait util the main goroutine ends even if it is not in WaitGroup.

	wg.Wait()
}

func channelSample() {
	ch := make(chan string)

	go outputByChannel("goroutine 1", 100, ch)
	go outputByChannel("goroutine 2", 300, ch)
	output("goroutine main", 200)

	// Wait until there are two "Done" pushed to the channel
	<-ch
	<-ch
}
