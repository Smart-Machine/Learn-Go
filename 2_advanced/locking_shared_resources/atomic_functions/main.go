package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	// counter is a variable incremented by all goroutines
	counter int64

	// wg is used to wait for the program to finish
	wg sync.WaitGroup
)

// The program is resolving race condition
// by the use of `atomic`, which permits
// the write to a goroutine at a time, but
// the downside of the atomic is the limitation
// to only the `int` data type.

func main() {
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()

	fmt.Println("Final Counter:", counter)
}

func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// Safely add one to counter
		atomic.AddInt64(&counter, 1)
		// Yield the thread and be placed back in the queue
		runtime.Gosched()
	}
}
