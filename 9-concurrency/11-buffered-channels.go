package main

// A send on an unbuffered channel can proceed if a receiver is ready.
// A send on a buffered channel can proceed if there is room in the buffer.
func main() {
	ch := make(chan int, 1)

	ch <- 100
	<-ch // receiving value from a buffered channel will remove the value from the buffer
	ch <- 200
}
