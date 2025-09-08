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

	// sender goroutine
	wg.Go(func() {
		x := slowFuncWithoutSelect(ctx)

		// unbuffered channel channel has one condition which is receiver is must
		// so if main goroutine moved on then there would be no receiver for the value
		// so using select we will only send the value if the receiver is ready
		// if ctx.Done happens then it means the receiver is not present anymore
		// and we should not send the value
		select {
		case <-ctx.Done():
			fmt.Println("sender: context is done, time up, and moving on")
			return
		case ch <- x:
			fmt.Println("sender: sent the value", x)
		}

	})

	//receiver goroutine (main)
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
	// if ctx.Done happens it means the receiver is not present anymore
	// so we are rolling back the changes to avoid an inconsistent state
	case <-ctx.Done():
		fmt.Println("rolling back all the changes as the context is done")
		return 0
	default:
	}

	// if ctx have time left then we are returning the value
	return 10
}

func slowFuncWithoutSelect(ctx context.Context) int {
	time.Sleep(2 * time.Second)
	fmt.Println("slowFunc() ran and changed 10 files")
	return 10
}
