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
	ch := make(chan string)

	for i := 0; i < 1000; i++ {
		go func(ch chan string) {
			totalThreadSafe.Mux.Lock()
			totalThreadSafe.Val++
			totalThreadSafe.Mux.Unlock()

			ch <- "Done"
		}(ch)

		<-ch
	}

	fmt.Println(totalThreadSafe.Val)
}
