package main

import (
	"fmt"
	"time"
)

func doWork(done <-chan bool) {
	i := 1
	for {
		select {
		case <-done:
			return
		default:
			for j := 0; j < i; j++ {
				fmt.Print("*")
			}
			fmt.Println()
			i++
		}
	}
}

func main() {
	done := make(chan bool)

	go doWork(done)

	time.Sleep(time.Millisecond * 1)

	close(done)
}
