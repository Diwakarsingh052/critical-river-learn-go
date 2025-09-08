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

var y = 1

func main() {
	wg := new(sync.WaitGroup)
	m := new(sync.RWMutex) // Read Write Mutex, RWMutex
	// using RWMutex we can take separate locks for reading and writing

	for i := 1; i <= 5; i++ {
		wg.Go(func() {
			UpdateY(i, m)
		})
		wg.Go(func() {
			PrintY(m)
		})
	}
	wg.Wait()

}

func UpdateY(val int, m *sync.RWMutex) {
	// critical section
	// this is the place where we access the shared resource
	func() {
		// when a goroutine acquires a lock, another goroutine can't access the critical section
		// until the lock is not released

		// when Write lock is acquired, no other read or writes are allowed
		m.Lock()
		defer m.Unlock() // release the lock when the function returns

		fmt.Println("Updating y variable takes some time")
		//time.Sleep(5 * time.Second)
		y = val
		fmt.Println("y variable updated")

		// here we work with x
	}()

	// no work with x
	//
	//
}

func PrintY(m *sync.RWMutex) {
	// Acquire a lock for reading
	m.RLock()
	// Releases the read lock when func completes

	//no one can write when read lock is acquired,
	// there could be unlimited number of reads
	defer m.RUnlock()
	fmt.Println("Printing y variable")
	fmt.Println(y)
}
