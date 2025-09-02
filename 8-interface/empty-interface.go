package main

import "fmt"

func main() {
	// research about type switch

	// empty interface , it can hold any type of data
	var i any // var i interface{} // both are same
	var a int
	i = 10

	// type assertion
	// using type assertion we can check the type of the interface and get the concrete value
	// i.(type) // type could be any valid concrete type
	a, ok := i.(int) // use type assertion always with two return values, value and ok
	// use ok variant to avoid panic
	if ok { // ok == true
		fmt.Println(a)
	}

	i = "hello"
	i = true
	i = 10.23
	i = []int{1, 2, 3}
	i = map[string]int{"a": 1, "b": 2}
	i = struct{ name string }{name: "bob"}

}
