package main

import "fmt"

func main() {
	a := []int{10, 20, 30, 40, 50}
	// slicing // referring to a part of the slice
	b := a[1:3] // index:len // // this line creates a new slice b
	fmt.Println(b)
	b = a[:3] // // start from 0 index till the length provided
	b = a[1:] // from the 1st index till the end
	fmt.Println(b)
}
