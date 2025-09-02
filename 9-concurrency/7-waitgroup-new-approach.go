package main

import (
	"fmt"
	"sync"
)

// this program works with go version 1.25 and above
func main() {
	// new returns a pointer to the type, type would be initialized to its default values
	wg := new(sync.WaitGroup)
	//wg := &sync.WaitGroup{}
	// anonymous goroutine function, function without name

	for i := 1; i <= 5; i++ {

		// wg.Go internally calls wg.Add(1) and defer wg.Done()
		wg.Go(func() {
			fmt.Println("work finished, id:", i)
		})
		wg.Go(func() {
			fmt.Println("add result", add(i, i))
		})

	}

	fmt.Println("main finished")
	wg.Wait()
}

func add(a, b int) int {
	return a + b
}
