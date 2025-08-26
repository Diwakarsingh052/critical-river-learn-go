package main

import "fmt"

func main() {
	var a [10]int // array //[0,0,0,0.....] // fixed size
	_ = a
	var x []int // slice // nil
	x[0] = 99   // this would panic, slice is nil ,
	// can't update an index which doesn't exist

	//a[10] = 100 //update
	a[9] = 200
	fmt.Println(a)

}
