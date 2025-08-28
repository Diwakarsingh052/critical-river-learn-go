package main

import "fmt"

func main() {
	type user struct {
		name string
		age  string
	}

	u := user{
		name: "bob",
		age:  "30",
	}

	fmt.Println(u)
	fmt.Printf("%T\n", u.name) // string , %T is for type
	fmt.Printf("%T\n", u)
}
