package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 10)

	wg := new(sync.WaitGroup)
	wgTasks := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
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
		// wait for all tasks to finish, then close the channel
		wgTasks.Wait()
		close(ch)
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
