package main

import "fmt"

func main() {
	fmt.Println(1, 2, 3, 4, 5)
	printArgs(10, 20, 30, 40, 50)
	printArgsV2([]int{1, 2, 3, 4, 5})
}

// in go using ... means that the function can take any number of arguments
// this is called variadic function
// we can't pass params after ...(variadic parameter)
// variadic parameter must be the last parameter
// variadic parameter is optional, if we don't pass any value then it will be nil'

func printArgs(i int, args ...int) {
	fmt.Printf("%T\n", args) // args is a slice
	fmt.Println(args)
}

func printArgsV2(args []int) {}
