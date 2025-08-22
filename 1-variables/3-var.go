package main

import "fmt"

func main() {
	x := 10
	{
		//shadowing
		fmt.Println(x)
		x := "hello"
		fmt.Println(x)
	}
	fmt.Println(x)
}

func abc() {
	x := 10
	func(a, b int) {
		//shadowing
		fmt.Println(x)
		x := "hello"
		fmt.Println(x)
	}(10, 20)
	fmt.Println(x)
}
