package main

import "fmt"

// Failed Update Mistake Number 2, Pointers changing references

// in Go everything is pass by value
// in case of pointers we copy the address from one part to another
var z int = 10

func main() {
	var p *int     // nil // default value of a pointer is nil
	p = &z         // x80
	updateValue(p) // copying x80 to p1 pointer

	fmt.Println(*p)

}

func updateValue(p1 *int) {
	var y int = 20 // x100

	// reference has been changed to a new address
	p1 = &y   // assigning x100 to p1
	*p1 = 200 // dereference and update the value
	fmt.Println(*p1)

}
