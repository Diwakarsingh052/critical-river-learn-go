package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//buffered()
	unbuffered()
}

func buffered() {
	// buffered channel doesn't care about the receiver
	// if buffer is empty it will put the value in the buffer and move on
	// a buffer channel allow more chance of parallelism
	// we need to manually make sure that all values are received

	wg := new(sync.WaitGroup)
	ch := make(chan int, 5) //[j6,j7,j8,j9,j10] j6,j7,j8,j9,j10

	wg.Go(func() {
		// sender goroutine
		for i := 1; i <= 5; i++ {
			ch <- i
			fmt.Println("sent value", i)
			fmt.Println("current number of values in buffered channel", len(ch))
		}
		fmt.Println("all values sent")
	})

	wg.Go(func() {
		for i := 0; i < 5; i++ {
			time.Sleep(1 * time.Second)
			fmt.Println(<-ch)
		}

	})

	wg.Wait()

}

func unbuffered() {

	// unbuffered channel will block the sender goroutine until the receiver is ready
	// unbufferd channel is a synchronous way of communication
	// good to create backpressure
	// it has guarantee that values must be received by a receiver
	wg := new(sync.WaitGroup)
	ch := make(chan int)
	wg.Go(func() {
		// sender goroutine
		for i := 1; i <= 5; i++ {
			ch <- i
			fmt.Println("sent value", i)
		}
		fmt.Println("all values sent")
	})

	wg.Go(func() {

		for i := 0; i < 5; i++ {
			//time.Sleep(1 * time.Second)
			fmt.Println(<-ch)
		}
	})

	wg.Wait()
}
