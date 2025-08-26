package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	greet()
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
	//markString := cmdArgs[2]

	fmt.Println(name)
	//var err error // default value of err is nil // which means no error
	// if there is an error there is a value in the form of a msg
	age, err := strconv.Atoi(ageString)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(age)

}
