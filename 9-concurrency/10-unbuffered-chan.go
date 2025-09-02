package main

import (
	"fmt"
	"sync"
)

// A send on an unbuffered channel can proceed if a receiver is ready.
func main() {
	wg := new(sync.WaitGroup)

	ch := make(chan int)

	wg.Go(func() {
		fmt.Println(<-ch) // blocking call
	})
	wg.Go(func() {
		ch <- 10
	})

	wg.Wait()

}
