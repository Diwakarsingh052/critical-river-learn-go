package main

import (
	"learn-go/3-slices/inspect"
)

func main() {
	x := []int{10, 20, 30, 40, 50, 60, 70}

	// using make we can preallocate an array
	// len is must when working with copy, if len is less the elems would be ignored
	b := make([]int, len(x), cap(x))
	copy(b, x) // deep copy, copying the actual data
	b[0] = 100
	inspect.InspectSlice("b", b)
	inspect.InspectSlice("x", x)

}
