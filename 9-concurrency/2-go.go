package main

import (
	"fmt"
	"time"
)

// concurrency is dealing with a lot of things at once
// parallelism is doing multiple things at once
func main() {
	// in go any func call or method call can be done in a goroutine or concurrently

	// below line doesn't mean that we are executing the goroutine
	go hello() // spinning up a new goroutine
	fmt.Println("hello from the main function")
	
	// worst way to wait for a goroutine to finish
	time.Sleep(time.Second * 1)
}

func hello() {
	fmt.Println("hello from the hello function")
}
