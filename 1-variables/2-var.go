package main

import "fmt"

var a int

const num = 2 // we can't use a function to give value to the constant

func main() {
	//var firstName string // camelCase for variable naming
	// every types have a default value
	var a int // int default is 0
	var b string = "hello"
	var c = "rahul"
	fmt.Println(a, b, c)
	//var d

	// go compiler would infer the type automatically from the right side value
	d, ok := 100, true // shorthand// create and assign
	_, _ = d, ok
	_, err := fmt.Println("hello") // assign values from the func call
	fmt.Println(err)
	//a = "1" // we cant change the type
	var x int
	var f float64 = float64(x)
	_ = f
}
