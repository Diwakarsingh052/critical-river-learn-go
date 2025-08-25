package main

import (
	"fmt"
	"learn-go/3-slices/inspect"
)

// If the capacity of s is not large enough to fit the additional values,
// append allocates a new, sufficiently large underlying array
// that fits both the existing slice elements and the additional values.
// Otherwise, append re-uses the underlying array.
func main() {
	x := []int{10, 20, 30, 40, 50, 60, 70}
	inspect.InspectSlice("x", x)
	// cap of would be calculated like
	// from the base index till the end of the backing array
	// in this case 5
	a := x[2:6]
	inspect.InspectSlice("a", a)

	fmt.Println("after append")
	//below line would change x, b refers to the same backing array,
	//there is room to add two more elements to the existing backing array
	//so it would add the value to the backing array refer  by x
	a = append(a, 888)

	//b = append(b, 777, 888, 999) // this would create a new backing array for b,
	//it will not change the x slice

	inspect.InspectSlice("a", a)
	inspect.InspectSlice("x", x)
}
