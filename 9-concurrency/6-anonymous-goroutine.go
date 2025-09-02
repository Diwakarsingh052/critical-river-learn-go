package main

import (
	"fmt"
	"sync"
)

func main() {
	// new returns a pointer to the type, type would be initialized to its default values
	wg := new(sync.WaitGroup)
	//wg := &sync.WaitGroup{}
	// anonymous goroutine function, function without name

	//wg.Add(5)
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Println("work finished, id:", id)
		}(i)
	} // () this is called as a function call

	fmt.Println("main finished")
	wg.Wait()
}
