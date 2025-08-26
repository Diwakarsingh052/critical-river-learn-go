package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	greet()
	fmt.Println("end of main")
}

func greet() {
	cmdArgs := os.Args[1:]
	// name, age, mark
	if len(cmdArgs) != 3 {
		log.Println("Invalid number of arguments")
		log.Println("Usage: ./greet <name> <age> <mark>")
		return // stops the current function execution
	}

	fmt.Println(cmdArgs)
	name := cmdArgs[0]
	ageString := cmdArgs[1]
	markString := cmdArgs[2]

	fmt.Println(name)
	//var err error // default value of err is nil // which means no error
	// if there is an error there is a value in the form of a msg

	//if you are calling a function , and if that func returns an error,
	//next thing MUST be error handling ,
	// you should not continue to write further logic
	age, err := strconv.Atoi(ageString)
	if err != nil {
		log.Println(err)
		//panic(err) // this will stop the execution of the program
		// os.Exit(1) // this will quit the program
		// log.Fatal(err) // it uses os.Exit(1) internally
		// not recommended to use // only during the app startup log.Fatal is fine
		return
	}
	marks, err := strconv.Atoi(markString)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(name, age, marks)

}
