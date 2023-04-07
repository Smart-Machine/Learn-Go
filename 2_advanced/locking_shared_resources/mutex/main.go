package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	// counter is a variable incremented by all goroutines
	counter int

	// wg is used to wait for the program to finish
	wg sync.WaitGroup

	// mutex is used to define a critical section of code
	mutex sync.Mutex
)

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
		// Only allow one goroutine through
		// this critical section at a time
		mutex.Lock()

		// Capture the value of counter
		value := counter

		// Yield the thread and be placed back in the queue
		runtime.Gosched()

		// Increment our local value of counter
		value++

		// Store the value back into counter
		counter = value

		// Release the lock and allow
		// any waiting goroutine through
		mutex.Unlock()

	}

}
