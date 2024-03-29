package main

import (
	"log"
	"sync"
	"time"
)

var names = []string{
	"ira",
	"mary",
	"nastea",
}

// namePrinter provides special support for printing names
type namePrinter struct {
	name string
}

// Task implements the Worker interface
func (m *namePrinter) Task() {
	log.Println(m.name)
	time.Sleep(time.Second * 4)
}

// Worker must be implemented by types that
// want to use the work pool
type Worker interface {
	Task()
}

// Pool provides a pool of goroutines that can execute any
// Worker tasks that are submitted
type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

// New creates a new work pool
func New(maxGoroutines int) *Pool {
	p := Pool{
		work: make(chan Worker),
	}

	p.wg.Add(maxGoroutines)
	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for w := range p.work {
				w.Task()
			}
			p.wg.Done()
		}()
	}

	return &p
}

// Run submits work to the pool
func (p *Pool) Run(w Worker) {
	p.work <- w
}

// Shutdown waits for all the goroutines to shutdown
func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}

func main() {
	// Create a work pool with 2 goroutines
	p := New(2)

	var wg sync.WaitGroup
	wg.Add(2 * len(names))

	for i := 0; i < 2; i++ {
		// Iterate over the slice of names
		for _, name := range names {
			// Create a namePrinter and
			// provide the specific name
			np := namePrinter{
				name: name,
			}

			go func() {
				// Submit the task to be worked on. When
				// RunTask return we know it is being handled
				p.Run(&np)
				wg.Done()
			}()
		}
	}

	wg.Wait()

	// Shutdown the work pool and wait for all existing work
	// to be completed
	p.Shutdown()
}
