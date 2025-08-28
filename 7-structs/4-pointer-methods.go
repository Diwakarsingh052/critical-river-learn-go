package main

import "fmt"

type author struct {
	name string
	age  int
}

// https://go.dev/doc/faq#methods_on_values_or_pointers
/*
First, and most important, does the method need to modify the receiver? If it does, the receiver must be a pointer.
Second is the consideration of efficiency. If the receiver is large, a big struct for instance, it may be cheaper to use a pointer receiver.
Next is consistency. If some of the methods of the type must have pointer receivers, the rest should too,
*/
func (a *author) updateName(name string) {
	a.name = name // no de-referencing needed
}

func main() {
	a := author{name: "zhangsan", age: 18}
	a.updateName("Bob") // address of a is passed to the method
	fmt.Println(a.name)

}
