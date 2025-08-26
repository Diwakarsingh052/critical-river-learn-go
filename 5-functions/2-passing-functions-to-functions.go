package main

import "fmt"

func main() {

	operate(add, 10, 20)
	operate(sub, 20, 30)
}

// operate func can accept function in op parameter,
// the function signature we are passing should match to op parameter type
func operate(op func(int, int), x, y int) {

	op(x, y)
}

// function signature is the function datatype
// datatype of func -> func(args)returnType
func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}
