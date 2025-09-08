package main

import (
	"fmt"
	"sync"
)

// BAD CODE AHEAD
// don't use infinite for select with buffered channels
// we can lose data
// if you want to use a buffered channel,
// then run goroutines and run range inside them to receive values
func main() {
	get := make(chan string, 1)
	post := make(chan string, 1)
	put := make(chan string, 1)
	done := make(chan struct{})

	wg := new(sync.WaitGroup)
	wgTasks := new(sync.WaitGroup)

	wgTasks.Go(func() {
		get <- "get"
	})

	wgTasks.Go(func() {
		post <- "post"
	})

	wgTasks.Go(func() {
		put <- "put"
		put <- "put 2"
	})

	wg.Go(func() {
		wgTasks.Wait()
		// close is also a send signal but without data
		close(done)
	})

	wg.Go(func() {
		for {
			select {
			//whichever case is not blocking exec that first
			//whichever case is ready first, exec that.
			// possible cases are chan recv , send , default
			case v := <-get:
				fmt.Println(v)
			case v := <-post:
				fmt.Println(v)
			case v := <-put:
				fmt.Println(v)
			case <-done:
				fmt.Println("all work done")
				return

			}
		}
	})

	wg.Wait()

}
