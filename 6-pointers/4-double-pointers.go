package main

import (
	"fmt"
)

// double pointers are useful when we want to update the original pointer
func main() {
	var p *int
	updateNilPointer(&p)
	fmt.Println(*p)

	// examples from standard library where double pointer is used
	//errors.As()
	//json.Unmarshal()
}

func updateNilPointer(p1 **int) {
	// assume p1 is storing x80(address of p from the main)
	x := 10 // assume address x90

	// trying to access the value of p1
	// which is also another pointer named as p from the main function
	*p1 = &x // updating x80 = x90 // it directly changes p from the main function itself
	//p = &x // above line means this
	**p1 = 200 // changing the value of x variable,
}
