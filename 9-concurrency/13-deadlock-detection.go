package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)
	ch := make(chan int)
	// go would not report deadlock if other goroutine is routine is running and not in blocked state
	wg.Go(func() {
		ch <- 10
	})
	wg.Go(func() {
		for {
			time.Sleep(1 * time.Second)
			fmt.Println("running forever")
		}
	})

	wg.Wait()
}
