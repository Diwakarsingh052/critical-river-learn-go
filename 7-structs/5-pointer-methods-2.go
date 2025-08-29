package main

import (
	"fmt"
)

type FileDetails struct {
	name string
}

func (f *FileDetails) PrintName() {

	fmt.Println(f.name)
}
func (f *FileDetails) UpdateName(fileName string) {

	f.name = fileName
}

func main() {
	var f1 FileDetails // x80{name:""}

	// try to avoid declaring struct variables as pointers,
	// it can cause nil pointer exceptions, if you forget to initialize the pointer
	//var f *FileDetails // nil

	f1.UpdateName("test.txt")
	f1.PrintName()

}
