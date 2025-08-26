package main

import (
	"fmt"
)

// using type we can define our own type or we can define alias
type money int // this is a new type

func main() {
	var i money = 100
	var j int = int(i)
	fmt.Println(j)

}
