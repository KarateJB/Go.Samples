package main

import (
	"fmt"
	"sync"
	"time"
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

	// Use unbuffered channel
	unbufChannelSample()

	// Use buffered Channel
	bufChannelSample()
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

	for i := 0; i < 1000; i++ {
		go func() {
			totalThreadSafe.Mux.Lock()
			totalThreadSafe.Val++
			totalThreadSafe.Mux.Unlock()
		}()
	}

	time.Sleep(1000 * time.Millisecond) // This block main goroutine to print the value of "totalSafe"
	fmt.Println(totalThreadSafe.Val)
}

func unbufChannelSample() {
	total := 0
	blockCh := make(chan string)
	for i := 0; i < 1000; i++ {
		go func(ch chan string) {
			total++
			blockCh <- "Done"
		}(blockCh)

		<-blockCh // Block each run in for loop until the goroutine in each run pushs
	}

	fmt.Println(total)
}

func bufChannelSample() {
	total := 0
	totalCh := make(chan int, 1) //Channel as int type and size is 1
	totalCh <- total

	for i := 0; i < 1000; i++ {
		go func() {

			totalCh <- (<-totalCh) + 1
		}()
	}

	time.Sleep(1000 * time.Millisecond) // This block main goroutine to print the value of "totalSafe"
	fmt.Println(<-totalCh)
	close(totalCh)
}
