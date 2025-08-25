package main

import "fmt"

func main() {
	x := []int{10, 20, 30}

	//UpdateSlice(x, 0, 100)
	x = UpdateAndAppend(x, 0, 100)
	fmt.Println(x)
}

func UpdateSlice(s []int, index, value int) {
	// if you want to update the value ,
	//then passing the reference of existing backing array is fine
	// we don't have to use copy
	s[index] = value
}

func UpdateAndAppend(s []int, index int, value int) []int {
	s[index] = value
	s = append(s, 1000)
	return s // make sure to return the updated slice if using append
	// or use double pointers
}
