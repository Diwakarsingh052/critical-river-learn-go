package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)
	wg := new(sync.WaitGroup)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	wg.Go(func() {
		x := slowFunc(ctx)
		select {
		case <-ctx.Done():
			fmt.Println("sender: context is done, time up, and moving on")
			return
		case ch <- x:
			fmt.Println("sender: sent the value", x)
		}

	})

	select {
	// if this case is true then it means the goroutine slowFunc() is not completed
	case <-ctx.Done():
		fmt.Println("context is done, time up, and moving on")

	case x := <-ch:
		// if this case is true then it means the goroutine slowFunc() is completed
		// and the value is received in the channel

		fmt.Println("received value", x)

	}
	wg.Wait()
}

func slowFunc(ctx context.Context) int {
	time.Sleep(2 * time.Second)
	fmt.Println("slowFunc() ran and changed 10 files")
	select {
	case <-ctx.Done():
		fmt.Println("rolling back all the changes as the context is done")
		return 0
	default:
	}
	return 10
}
