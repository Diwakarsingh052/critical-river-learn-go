package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int, 10)

	wg := new(sync.WaitGroup)
	wgTasks := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		time.Sleep(1 * time.Second)
		defer wg.Done()
		for i := 1; i <= 5; i++ {
			wgTasks.Add(1)
			// fan out pattern
			// spinning n number of goroutines for n tasks
			go func() {
				defer wgTasks.Done()
				ch <- i // send
			}()
		}

		// running another goroutine to wait for all tasks to finish
		wg.Go(func() {
			// wait for all tasks to finish, then close the channel
			wgTasks.Wait()
			close(ch)
		})

	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range ch {
			fmt.Println("received", v)
		}
	}()

	wg.Wait()
}
