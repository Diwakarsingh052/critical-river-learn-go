package main

import "fmt"

type Speaker interface {
	Speak() string
	Speak2() string
}

type Person struct {
	name string
}

func (p *Person) Speak() string {
	return fmt.Sprintf("Hi, my name is %s", p.name)
}
func (p Person) Speak2() string {
	return fmt.Sprintf("Hi, my name is 2 %s", p.name)
}

func main() {
	// if method set is same as of the interface methods
	// then only interface is implemented by a type variable
	p := Person{
		name: "james",
	}

	// method set of p of person type
	// Speak2()

	var s Speaker = &p // as we are passing a pointer to the interface
	// method set would now include both Speak() and Speak2() methods
	fmt.Println(s.Speak())

	p1 := &Person{
		name: "james",
	}
	// p1 method set, p1 can associate with both value and pointer pointer methods
	// Speak(), Speak2()
	var s1 Speaker = p1

	fmt.Println(s1.Speak())
}

/*
                          +-------------------------------------+
                          |              Method Set             |
                          +------------------+------------------+
                          |    Value Type    |   Pointer Type   |
+-------------------------+------------------+------------------+
| Function w/ Value Rec.  |        Yes       |       Yes        |
| ( func (t T) )          |                  |                  |
+-------------------------+------------------+------------------+
| Function w/ Ptr Rec.    |        No        |       Yes        |
| ( func (t *T) )         |                  |                  |
+-------------------------+------------------+------------------+

If a function is implemented with a value receiver (func (t T)),
it can be called through a value or a pointer.

If a function is implemented with a pointer receiver (func (t *T)),
it can only be called through a pointer.
*/
