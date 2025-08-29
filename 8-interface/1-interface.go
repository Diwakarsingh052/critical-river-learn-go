package main

import (
	"fmt"
	"io"
)

// Polymorphism means that a piece of code changes its behavior depending on the
// concrete data it’s operating on // Tom Kurtz, Basic inventor

// interface is an abstract type

// interfaces are implemented by concrete types automatically, if a struct has
// all the methods required by the interface

// "Don’t design with interfaces, discover them". - Rob Pike

// Bigger the interface weaker the abstraction // Rob Pike

//	type Reader interface {
//		Read(b []byte) (int, error)
//		//abc()  // all methods must be implemented to implement the interface
//	}
type File struct {
	Name string
}

func (f File) Read(b []byte) (int, error) {
	fmt.Println("reading files and processing them", f.Name)
	return 0, nil
}

type IO struct {
	name string
}

func (i IO) Read(b []byte) (int, error) {
	fmt.Println("reading and processing io ", i.name)
	return 0, nil
}

//	func DoReading(f File, i IO) {
//		f.Read(nil)
//		i.Read(nil)
//	}

// r is a variable of an abstract type, which is an interface

func DoReading(r io.Reader) {
	r.Read(nil)
	fmt.Printf("%T\n", r)
}

func main() {
	f := File{Name: "abc"}
	i := IO{name: "io"}
	DoReading(f)
	DoReading(i)

}
