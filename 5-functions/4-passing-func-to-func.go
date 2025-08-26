package main

import "fmt"

func main() {

	operateV3(addV2(), 10, 20)
	////fmt.Println(add)
	//fmt.Println(add(10, 20))

}

// operate func can accept function in op parameter,
// the function signature we are passing should match to op parameter type
func operateV3(op func(int, int) int, x, y int) {
	sum := op(x, y)
	fmt.Println(sum)
}

// function signature is the function datatype
// datatype of func -> func(args)returnType
func addV2() func(int, int) int {

	// in GO we can store function body in variable
	f := func(a, b int) int {
		return a + b
	}
	return f
}

func subV2(a, b int) int {
	return a - b
}
