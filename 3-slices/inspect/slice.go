package inspect

import "fmt"

func InspectSlice(name string, slice []int) {
	fmt.Printf("name %v Length %d Cap %d \n", name, len(slice), cap(slice))
	fmt.Printf("memory address %p\n", slice) // memory address of the first element
	fmt.Println(slice)
	fmt.Println()

}
