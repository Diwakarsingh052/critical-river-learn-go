package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func OpenFile(name string) (*os.File, error) {
	f, err := os.Open(name)
	if err != nil {

		//errors.Is can check if an error was wrapped inside the chain or not
		//  if an error was found in the chain, you now know what exactly went wrong
		// you might want to take some actions to fix the issue
		//or maybe just log the additional details
		if errors.Is(err, os.ErrNotExist) {
			log.Println("file not found, but trying to create a new one")

			// create a new file
			f, err := os.Create(name)
			if err != nil {
				// if it still fails, we will return the error
				return nil, err
			}

			// success case // errors.Is succeeded in identifying the root cause of the problems
			return f, nil
		}
		// file opening failed for some other reason, return the error
		return nil, err
	}
	// file opened successfully
	return f, nil
}

func main() {
	f, err := OpenFile("file.txt")
	if err != nil {
		log.Println(err)
		return
	}
	info, _ := f.Stat()
	fmt.Println(info.Name())

}
