package main

import (
	"fmt"
)

type Student struct {
	name string
	age  int
}

// method signature
// func (receiver) funcSignature {//body}

func (s Student) Print() {
	fmt.Println(s.name)
}

func PrintName(s Student) {
	fmt.Println(s.name)
}

func main() {
	s := Student{
		name: "bob",
		age:  30,
	}
	PrintName(s)
	// methods could be only called using struct variable
	s.Print()

}
