package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Open("file.txt")
	// this line would exec when the main function returns or panic
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

}
