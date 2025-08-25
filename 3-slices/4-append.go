package main

import (
	"fmt"
	"learn-go/3-slices/inspect"
)

// go through the link for append working
// https://go.dev/ref/spec#Appending_and_copying_slices
func main() {
	var s []int
	inspect.InspectSlice("s", s)

	// whenever shorthand is used , len and cap are equal
	x := []int{10, 20, 30, 40, 50}
	inspect.InspectSlice("x", x)

	a := make([]int, 0, 10)
	// with make we can plan out number of elements to be stored
	//in slice before hand, it helps in optimizing
	//the memory allocation and deallocation part

	inspect.InspectSlice("a", a)

	fmt.Println("after")
	s = append(s, 10, 20, 30, 40, 50)
	inspect.InspectSlice("s", s)

	x = append(x, 60, 70)
	inspect.InspectSlice("x", x)

	a = append(a, 10, 20)
	inspect.InspectSlice("a", a)
	
	fmt.Println(s)
	fmt.Println(x)
}
