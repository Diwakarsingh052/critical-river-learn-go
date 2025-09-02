package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)

	wg.Go(func() {
		// we can't return values from goroutine directly to other goroutine
		id := userId()
		fmt.Println(id)
	})

	fmt.Println("main finished ")
	
	wg.Wait()

}

func userId() int {
	return 101
}
