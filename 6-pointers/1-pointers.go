package main

import "fmt"

func main() {
	// default value of pointer is nil
	//var p *int // pointer to int
	var a int = 10
	//p = &a
	// print the address of a
	fmt.Println(&a) // & is the address operator

	// passing the address of a
	updateWithPointer(&a)
	fmt.Println(a)
}

func updateWithoutPointer(a int) int {
	a = 100
	return a
}

// pointer holds the memory address of a variable
func updateWithPointer(p *int) {
	fmt.Println(*p) // print the value at the address
	// * is the dereference operator
	// * means value at the address

	// below line is trying to update the address of pointer which is not allowed
	//p = 200 // this line is not valid, we can't assign a normal value to a pointer'
}
