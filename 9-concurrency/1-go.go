package main

import (
	"fmt"
	"runtime"
)

func main() {
	//set number of cores to 1, returns the previous setting of cpu
	fmt.Println(runtime.GOMAXPROCS(9))
	fmt.Println(runtime.GOMAXPROCS(runtime.NumCPU()))
	// panic reveals goroutine id
	//panic("hello")
}
