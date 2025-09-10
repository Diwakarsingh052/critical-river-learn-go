package main

import (
	"fmt"
	"runtime/debug"
)

// panic is a runtime exception
// if the caller function doesn't depends on the called function you can stop panic propagation back
// by calling the recovery function in defer block

// defer guarantees to run // so it would recover the panic if it would happen
func main() {

	DoSomething()
	fmt.Println("main doing further stuff")
	fmt.Println("work doesnt depends on the result of DoSomething")
	fmt.Println("end of main")
}

func DoSomething() {
	// we need to decide where we would recover from the panic
	// the func where we decide to recover the panic needs to stop

	// RecoverPanic would recover the current function from panic, but the function needs to stop
	// it can't continue executing
	defer recoverPanic()
	updateSlice(nil)
	fmt.Println("end of DoSomething")
}

func updateSlice(s []int) {

	s[0] = 100
	fmt.Println("end of the updateSlice")
}

func recoverPanic() {
	// msg is of the any type
	// default value of any type is nil
	msg := recover()

	if msg != nil {
		//	 if msg is not nil then panic occured
		fmt.Println("recovered from the panic")
		fmt.Println(msg)
		fmt.Println(string(debug.Stack()))
	}
}
