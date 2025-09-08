package main

import (
	"fmt"
	"sync"
)

// shared resources
/*
1. Global Variables
2. Pointer variables
3. Structs fields where methods are pointers
4. Maps,
*/

// data race situations
//	- at least one concurrent write and n number of reads
//	- n number of concurrent writes
// 	- n number of concurrent writes and n number of concurrent reads
// 	Note - Data race doesn't happen if there are only concurrent reads

// run program with race detector
// go run -race main.go // don't use this in production
var x = 1

func main() {
	wg := new(sync.WaitGroup)
	m := new(sync.Mutex)

	for i := 1; i <= 5; i++ {
		wg.Go(func() {
			UpdateX(i, m)
		})
		wg.Go(func() {
			PrintX(m)
		})
	}
	wg.Wait()

}

func UpdateX(val int, m *sync.Mutex) {
	// critical section
	// this is the place where we access the shared resource
	func() {
		// when a goroutine acquires a lock, another goroutine can't access the critical section
		// until the lock is not released
		m.Lock()
		defer m.Unlock() // release the lock when the function returns
		x = val
		// here we work with x
	}()

	// no work with x
	//
	//
}

func PrintX(m *sync.Mutex) {
	m.Lock()
	defer m.Unlock()
	fmt.Println(x)
}
