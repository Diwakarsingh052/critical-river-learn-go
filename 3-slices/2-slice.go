package main

import "fmt"

func main() {

	a := []int{10, 20, 30}
	//var b []int = []int{10, 20, 30}

	//for i := 0; i < len(a)-1; i++ {
	//	fmt.Println(a[i])
	//}

	for i, v := range a { // range returns two values, index and value
		fmt.Println(i, v)
	}

}
