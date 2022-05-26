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
	waitGroupSample()

	// Channel
}

func output(s string, delay int) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Duration(delay) * time.Millisecond)
		fmt.Println(s)
	}
}

func outputWg(s string, delay int, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done() // decrease counter by 1, once counter eauals 0, WaitGroup stop blocking.
	}

	for i := 0; i < 5; i++ {
		time.Sleep(time.Duration(delay) * time.Millisecond)
		fmt.Println(s)
	}
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
	go outputWg("goroutine 1", 100, wg)
	go outputWg("goroutine 2", 300, wg)
	// output("goroutine main", 200, nil) // The process will still wait util the main goroutine ends even if it is not in WaitGroup.
	wg.Wait()

	outputWg("goroutine main", 200, nil)
}
