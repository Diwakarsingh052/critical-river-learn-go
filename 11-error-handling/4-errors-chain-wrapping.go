package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

var ErrFileNotFound = errors.New("file not found")

func openFile() (*os.File, error) {
	f, err := os.Open("file2.txt")
	if err != nil {
		// using fmt.Errorf to wrap the error
		// we are creating a chain of errors
		// %w is used to wrap the error
		return nil, fmt.Errorf("%w %w", err, ErrFileNotFound)
	}
	return f, nil

}

func main() {
	_, err := openFile()
	if err != nil {
		if errors.Is(err, ErrFileNotFound) {
			fmt.Println("custom error found in the chain")
			log.Println(err)
			return
		}
		log.Println(err)
		return
	}
	fmt.Println("file opened successfully")
}
