package main

import (
	"fmt"
	"sync"
)

// ThreadSafeNumber
type ThreadSafeNumber struct {
	Val int
	Mux sync.Mutex // mutex
}

func main() {
	// NG
	ngSample()

	// Use Mutex
	mutexSample()

	// Use Channel
	channelSample()
}

func ngSample() {
	totalNonThreadSafe := 0
	for i := 0; i < 1000; i++ {
		go func() {
			totalNonThreadSafe++
		}()
	}

	fmt.Println(totalNonThreadSafe)
}

func mutexSample() {
	totalThreadSafe := ThreadSafeNumber{Val: 0}
	blockCh := make(chan string)

	for i := 0; i < 1000; i++ {
		go func(ch chan string) {
			totalThreadSafe.Mux.Lock()
			totalThreadSafe.Val++
			totalThreadSafe.Mux.Unlock()

			ch <- "Done"
		}(blockCh)

		<-blockCh // receiving value from channel
	}

	fmt.Println(totalThreadSafe.Val)
	close(blockCh)
}

func channelSample() {
	total := 0
	blockCh := make(chan string)
	totalCh := make(chan int, 1)
	totalCh <- total

	for i := 0; i < 1000; i++ {
		go func(ch chan string) {

			totalCh <- (<-totalCh) + 1

			ch <- "Done"

		}(blockCh)

		<-blockCh
	}

	fmt.Println(<-totalCh)
	close(blockCh)
	close(totalCh)
}
