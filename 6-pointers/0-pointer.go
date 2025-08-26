package main

import "fmt"

func main() {
	var p *int
	var x int = 10
	p = &x
	fmt.Println(x)
	fmt.Println(&x)
	fmt.Println(p)
	fmt.Println(*p)
}
