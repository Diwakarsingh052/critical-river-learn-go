package main

import (
	"fmt"
	"sync"
	"time"
)

// A send on an unbuffered channel can proceed if a receiver is ready.
// A send on a buffered channel can proceed if there is room in the buffer.
func main() {
	wg := new(sync.WaitGroup)
	// creating a channel of type int
	ch := make(chan int)

	wg.Go(func() {
		// wg.Add and wg.Done are done automatically by wg.GO
		fmt.Println("working on  work id", 10)
		time.Sleep(5 * time.Second)
		ch <- 10 // send value to channel
	})

	// receiver would block until the value is sent
	id := <-ch // receive value from channel

	fmt.Println("received id", id)
	fmt.Println("main done")
	wg.Wait()

}
