package main

import (
	"fmt"
	"sync"
	"time"
)

// use this for go version before 1.25
func main() {
	wg := &sync.WaitGroup{}
	// waitgroup counter represents number of goroutine we are running
	wg.Add(1) // adding 1 to the counter
	go helloV2(wg)

	fmt.Println("more work going on in the main function")

	// go scheduler will schedule another goroutine to run if wait blocks
	wg.Wait() // wait until the counter is not back to 0
	fmt.Println("main done")

}

func helloV2(wg *sync.WaitGroup) {

	defer wg.Done() // decrement the counter by one
	time.Sleep(time.Second * 2)
	fmt.Println("hello")

}
