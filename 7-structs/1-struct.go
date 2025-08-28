package main

import "fmt"

type Person struct {
	name    string // fields of struct
	age     int
	address string
}

func main() {

	//var loginDetails struct {
	//	username string
	//	password string
	//}

	var p Person // p is a struct variable
	p.name = "bob"

	fmt.Printf("%+v", p) // it can print field value pairs
}
